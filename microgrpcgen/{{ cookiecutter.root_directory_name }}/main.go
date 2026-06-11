package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	{% if cookiecutter.use_database == 'y' %}db "{{ cookiecutter.module_name }}/db"{% endif %}
	{% if cookiecutter.use_database == 'y' %}"github.com/jmoiron/sqlx"{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}kafka "{{ cookiecutter.module_name }}/internal/kafka"{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}"github.com/sanservices/kit/kafkalistener"{% endif %}
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.module_name }}-proto/pb"
	"google.golang.org/grpc/reflection"
	config "{{ cookiecutter.module_name }}/config"
	api "{{ cookiecutter.module_name }}/internal/api"
	handler "{{ cookiecutter.module_name }}/internal/api/v1"
	healthcheck "{{ cookiecutter.module_name }}/internal/api/healthcheck"
	repository "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository"
	{% if cookiecutter.use_cache == 'y' %}redis "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository/redis"{% endif %}
	service "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/service"
	{% if cookiecutter.use_cache == 'y' %}{{ cookiecutter.service_name }} "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"{% endif %}
	"github.com/labstack/echo/v4"
	log "github.com/sanservices/apilogger/v2"
	ddtracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			// Initialize context
			context.Background,
			
			// Initialize log
			log.New,
			
			// Initialize service configuration
			config.New,
			
			{% if cookiecutter.use_database == 'y' %}
			// Initialize database connection
			db.New,
			{% endif %}

			{% if cookiecutter.use_cache == 'y' %}
			// Initialize redis connection (exposed as the {{ cookiecutter.service_name }}.Cache interface so fx can inject it)
			fx.Annotate(
				redis.New,
				fx.As(new({{ cookiecutter.service_name }}.Cache)),
			),
			{% endif %}

			// Initialize repository layer for databases transactions
			repository.New,

			// Initialize service layer for business logic
			service.New,

			{% if cookiecutter.use_kafka == 'y' %}
			//Initialize kafka's message broker
			kafkalistener.New,
			
			// Initialize kafka implementation
			kafka.New,
			{% endif %}

			// Initialize api server
			api.New,
			handler.New,
			healthcheck.New,

			// Assemble the echo handlers (healthcheck + swagger docs). handler.New
			// is reused here and for the gRPC server, so the same instance serves
			// both surfaces.
			func(h *handler.Handler, hc *healthcheck.Healthcheck) []api.Handler {
				return []api.Handler{
					hc,
					h,
				}
			},
		),

		fx.Invoke(
			// Print log startup
			func(ctx context.Context, l *log.Logger) {
				l.Info(ctx, log.LogCatStartUp, "Initializing {{ cookiecutter.root_directory_name }} service")
			},

			// Apply middleware and register the echo handlers (healthcheck + swagger docs)
			api.RegisterRoutes,

			// Adds the OnStart & OnStop callbacks
			func(lc fx.Lifecycle, ctx context.Context, config *config.Settings, handler *handler.Handler, e *echo.Echo, {% if cookiecutter.use_database == 'y' %}db *sqlx.DB,{% endif %} {% if cookiecutter.use_kafka == 'y' %}k *kafka.Kafka{% endif %}) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						// Start the Datadog tracer with basic service info. The agent
						// address and other options come from the standard DD_* env vars.
						ddtracer.Start(
							ddtracer.WithService(config.Service.Name),
							ddtracer.WithServiceVersion(config.Service.Version),
							ddtracer.WithEnv(os.Getenv("DD_ENV")),
						)

						// Start the gRPC server synchronously so a bind failure aborts startup.
						if err := StartGRPCServer(ctx, config, handler); err != nil {
							return err
						}

						go startRestAPI(ctx, config, e)
						{%- if cookiecutter.use_kafka == 'y' %}
						go k.StartListener(ctx)
						{% endif %}

						return nil
					},

					OnStop: func(ctx context.Context) error {
						{%- if cookiecutter.use_database == 'y' %}
						log.Info(ctx, log.LogCatDatastoreClose, "closing database...")
						if err := db.Close(); err != nil {
							log.Errorf(ctx, log.LogCatDatastoreClose, "error closing database: %v", err)
							return err
						}
						{% endif %}
						log.Info(ctx, log.LogCatUncategorized, "server is shutting down...")
						if err := e.Shutdown(ctx); err != nil {
							log.Errorf(ctx, log.LogCatUncategorized, "error shutting down server: %v", err)
							return err
						}

						// Flush and stop the Datadog tracer.
						ddtracer.Stop()

						return nil
					},
				})
			},
		),
	)
	app.Run()
}

func startRestAPI(ctx context.Context, config *config.Settings, e *echo.Echo) {

	// gRPC-gateway catch-all. The static routes registered by api.RegisterRoutes
	// (healthcheck, /v1/docs) take precedence over this in echo's router.
	e.Any("/*", echo.WrapHandler(setupGrpcGatewayHandler(ctx, config)))

	address := fmt.Sprintf(":%d", config.Service.Port)
	log.Infof(ctx, log.LogCatStartUp, "starting REST API on port %d (swagger at http://localhost:%d/v1/docs)", config.Service.Port, config.Service.Port)

	// http.ErrServerClosed is returned on a graceful shutdown and is expected.
	if err := e.Start(address); err != nil && err != http.ErrServerClosed {
		log.Errorf(ctx, log.LogCatUncategorized, "REST API server failed: %v", err)
	}
}

func setupGrpcGatewayHandler(ctx context.Context, config *config.Settings) http.Handler {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	address := fmt.Sprintf("0.0.0.0:%s", config.GRPC.Port)
	err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, address, opts)
	if err != nil {
		log.Errorf(ctx, log.LogCatUncategorized, "failed to register gRPC-Gateway: %v", err)
	}

	return mux
}

func StartGRPCServer(ctx context.Context, config *config.Settings, handler *handler.Handler) error {

	log.Infof(ctx, log.LogCatStartUp, "starting gRPC server on port %s", config.GRPC.Port)
	address := fmt.Sprintf("0.0.0.0:%s", config.GRPC.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf(ctx, log.LogCatStartUp, "failed to listen on %s: %v", address, err)
		return err
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterUserServer(grpcServer, handler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Errorf(ctx, log.LogCatUncategorized, "gRPC server failed: %v", err)
		}
	}()

	return nil
}
