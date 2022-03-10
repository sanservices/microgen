package fxhooks

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"{{ cookiecutter.module_name }}/goutils/settings"
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
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go startRestAPI(config, e)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := e.Shutdown(ctx)
			if err != nil {
				log.Println("Error shutting down echo server:", err)
			}

			err = db.Close()
			if err != nil {
				log.Println("Error closing database connection:", err)
			}

			return nil
		},
	})
}
