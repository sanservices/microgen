package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	{% if cookiecutter.use_database == 'y' %}db "{{ cookiecutter.module_name }}/db"{% endif %}
	{% if cookiecutter.use_database == 'y' %}"github.com/jmoiron/sqlx"{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}kafka "{{ cookiecutter.module_name }}/internal/kafka"{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}"github.com/sanservices/kit/kafkalistener"{% endif %}
	config "{{ cookiecutter.module_name }}/config"
	api "{{ cookiecutter.module_name }}/internal/api"
	handler "{{ cookiecutter.module_name }}/internal/api/v1"
	healthcheck "{{ cookiecutter.module_name }}/internal/api/healthcheck"
	{{ cookiecutter.service_name }} "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
	repository "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository"
	{% if cookiecutter.use_cache == 'y' %}redis "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository/redis"{% endif %}
	service "{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/service"
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
			
			// Initialize handlers
			func(config *config.Settings, service {{ cookiecutter.service_name }}.Service) []api.Handler {
				return []api.Handler{
					healthcheck.New(),
					handler.New(config, service),
				}
			},
		),

		fx.Invoke(
			// Print log startup
			func(ctx context.Context, l *log.Logger) {
				l.Info(ctx, log.LogCatStartUp, "Initializing {{ cookiecutter.root_directory_name }} service")
			},

			// Enables the REST API endpoints
			api.RegisterRoutes,
			
			// Adds the OnStart & OnStop callbacks
			func(lc fx.Lifecycle, ctx context.Context, config *config.Settings, e *echo.Echo, {% if cookiecutter.use_database == 'y' %}db *sqlx.DB,{% endif %} {% if cookiecutter.use_kafka == 'y' %}k *kafka.Kafka{% endif %}) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						// Start the Datadog tracer with basic service info. The agent
						// address and other options come from the standard DD_* env vars.
						ddtracer.Start(
							ddtracer.WithService(config.Service.Name),
							ddtracer.WithServiceVersion(config.Service.Version),
							ddtracer.WithEnv(os.Getenv("DD_ENV")),
						)

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
	address := fmt.Sprintf(":%d", config.Service.Port)

	log.Infof(ctx, log.LogCatStartUp, "starting REST API on port %d (swagger at http://localhost:%d/v1/docs)", config.Service.Port, config.Service.Port)

	// http.ErrServerClosed is returned on a graceful shutdown and is expected.
	if err := e.Start(address); err != nil && err != http.ErrServerClosed {
		log.Errorf(ctx, log.LogCatUncategorized, "REST API server failed: %v", err)
	}
}
