// {{cookiecutter.app_name}} service
//
// {{cookiecutter.app_name}} service
//
// Schemes: http
// Host: localhost:8080
// BasePath: /v1
// Version: 1.0
//
// Security:
//     - api_key:
//
// SecurityDefinitions:
//  api_key:
//   type: apiKey
//   name: api-key
//   in: header
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	logger "github.com/sanservices/apilogger/v2"
	"go.uber.org/fx"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}/repository"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}/service"
	"{{cookiecutter.module_name}}/internal/api"
	"{{cookiecutter.module_name}}/internal/api/healthcheck"
	"{{cookiecutter.module_name}}/internal/api/v1"
	"{{cookiecutter.module_name}}/settings"
)

func main() {
	fx.New(
		// Set logger to custom one
		fx.Logger(logger.New()),

		fx.Provide(
			// Provide new instances of structs
			// Empty context
			context.Background,
			// Logger
			logger.New,
			// Settings
			settings.New,
			// Repo
			repository.New,
			// Service
			service.New,
			// Validator
			api.NewValidator,
			// New server
			api.NewServer,
			// Add all handlers here
			func(cfg *settings.Settings, srv {{cookiecutter.main_domain}}.Service, vld *api.Validator) []api.Handler {
				return []api.Handler{
					healthcheck.NewHandler(),     // For Healthchecks
					v1.NewHandler(cfg, srv, vld), // v1
				}
			},
		),
		fx.Invoke(
			// Use logger to initialize it globally
			func(ctx context.Context, l *logger.Logger) {
				logger.Info(ctx, logger.LogCatStartUp, "Initializing the app")
			},

			// Register routes
			api.RegisterRoutes,
		),
	).Run()
}
