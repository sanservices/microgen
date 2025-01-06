package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
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
	"github.com/labstack/echo/v4"
	log "github.com/sanservices/apilogger/v2"
	echoMW "github.com/labstack/echo/v4/middleware"
	apicoreMW "github.com/sanservices/apicore/middleware"
	log "github.com/sanservices/apilogger/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			// Initialize context
			context.Background,
			
			// Intialize log
			log.New,
			
			// Intialize service configuration
			config.New,
			
			{% if cookiecutter.use_database == 'y' %}
			// Intialize database connection
			db.New,
			{% endif %}

			{% if cookiecutter.use_cache == 'y' %}
			// Initialize redis connection
			redis.New,
			{% endif %}

			// Intialize repository layer for databases transactions
			repository.New,

			// Intialize service layer for buisness logic
			service.New,

			{% if cookiecutter.use_kafka == 'y' %}
			//Initialize kafka's message broker
			kafkalistener.New,
			
			// Initialize kafka implemetation
			kafka.New,
			{% endif %}

			// Intialize api server
			api.New,
			handler.New,
			healthcheck.New,
		),

		fx.Invoke(
			// Print log startup
			func(ctx context.Context, l *log.Logger) {
				l.Info(ctx, log.LogCatStartUp, "Initializing {{ cookiecutter.root_directory_name }} service")
			},


			// Adds the OnStart & OnStop callbacks
			func(lc fx.Lifecycle, ctx context.Context, config *config.Settings, handler *handler.Handler, e *echo.Echo, healthcheck *healthcheck.Healthcheck, {% if cookiecutter.use_database == 'y' %}db *sqlx.DB,{% endif %} {% if cookiecutter.use_kafka == 'y' %}k *kafka.Kafka{% endif %}) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go StartGRPCServer(config, handler)
						go startRestAPI(ctx, config, e, healthcheck)
						{%- if cookiecutter.use_kafka == 'y' %}
						go k.StartListener(ctx)
						{% endif %}
						
						return nil
					},

					OnStop: func(ctx context.Context) error {
						{%- if cookiecutter.use_database == 'y' %}
						log.Info(ctx, log.LogCatDatastoreClose, "Closing database...")
						if err := db.Close(); err != nil {
							log.Error(ctx, log.LogCatDatastoreClose, "Error closing database")
							return err
						}
						{% endif %}
						log.Info(ctx, log.LogCatUncategorized, "Server is shutting down...")
						if err := e.Shutdown(ctx); err != nil {
							log.Error(ctx, log.LogCatUncategorized, "Error shutting down server")
							return err
						}

						return nil
					},
				})
			},
		),
	)
	app.Run()
}

func startRestAPI(ctx context.Context, config *config.Settings, e *echo.Echo, healthcheck *healthcheck.Healthcheck) {

	healthcheck.RegisterRoutes(e.Group(""))
	e.Use(apicoreMW.SetCustomHeaders)
	e.Use(apicoreMW.EnrichContext)
	e.Use(apicoreMW.RequestLogger)
	e.Use(echoMW.Recover())
	e.Any("/*", echo.WrapHandler(setupGrpcGatewayHandler(config)))
	address := fmt.Sprintf(":%d", config.Service.Port)
	log.Infof(ctx, log.LogCatUncategorized, "See swagger at http://localhost:%d/v1/docs", config.Service.Port)

	e.Logger.Fatal(e.Start(address))
}

func setupGrpcGatewayHandler(config *config.Settings) http.Handler {

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	address := fmt.Sprintf("0.0.0.0:%s", config.GRPC.Port)
	err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, address, opts)
	if err != nil {
		log.Errorf(ctx, log.LogCatUncategorized, "failed to register gRPC-Gateway: %v", err)
	}

	return mux
}

func StartGRPCServer(config *config.Settings, handler *handler.Handler) error {

	ctx := context.Background()
	log.Infof(ctx, log.LogCatUncategorized, "initiating rpc server on port:%s", config.GRPC.Port)
	address := fmt.Sprintf("0.0.0.0:%s", config.GRPC.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error(ctx, log.LogCatUncategorized, err)
		return err
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterUserServer(grpcServer, handler)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Error(ctx, log.LogCatUncategorized, fmt.Sprintf("gRPC server failed: %v", err))
		}
	}()
	return nil
}
