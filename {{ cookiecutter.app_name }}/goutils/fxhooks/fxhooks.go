package fxhooks

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"{{ cookiecutter.module_name }}/goutils/settings"
	{% if cookiecutter.include_kafka == 'Yes' %}
	"{{ cookiecutter.module_name }}/internal/kafka"
	{% endif %}
	"go.uber.org/fx"
)

func startRestAPI(config *settings.Settings, e *echo.Echo) {
	address := fmt.Sprintf(":%d", config.Service.Port)

	log.Printf("See swagger at http://localhost:%d/v1/docs", config.Service.Port)
	err := e.Start(address)
	if err != nil {
		log.Println(err)
	}
}

func SetLifeCycleHooks(
	config *settings.Settings,
	lc fx.Lifecycle,
	e *echo.Echo,
	db *sqlx.DB,
	k *kafka.Kafka,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go startRestAPI(config, e)
			{% if cookiecutter.include_kafka == 'Yes' %}
			go k.StartListener(ctx)
			{% endif %}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := e.Shutdown(ctx)
			if err != nil {
				log.Println("Error shutting down echo server:", err)
			}

			{% if cookiecutter.include_kafka == 'Yes' %}
			err = k.StopListener()
			if err != nil {
				log.Println("Error stopping kafka listener:", err)
			}
			{% endif %}

			err = db.Close()
			if err != nil {
				log.Println("Error closing database connection:", err)
			}

			return nil
		},
	})
}
