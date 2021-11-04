// {{ cookiecutter.app_name }} service
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

	"{{ cookiecutter.module_name }}/goutils/dbfactory"
	"{{ cookiecutter.module_name }}/goutils/settings"
	"{{ cookiecutter.module_name }}/internal/api"
	"{{ cookiecutter.module_name }}/internal/api/healthcheck"
	v1 "{{ cookiecutter.module_name }}/internal/api/v1"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/service"
	"github.com/sanservices/apicore/validator"
	logger "github.com/sanservices/apilogger/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Set logger to custom one
		// fx.Logger(logger.New()),

		fx.Provide(
			// Provide new instances of structs
			// Empty context
			context.Background,
			// Logger
			logger.New,
			// Settings
			settings.New,
			// database connection
			dbfactory.New,
			// Repo
			repository.New,
			// Service
			service.New,
			// Validator
			validator.NewValidator,
			// New server
			api.NewServer,
			// Add all handlers here
			func(cfg *settings.Settings, srv app.Service, vld *validator.Validator) []api.Handler {
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
	)

	app.Run()
}
