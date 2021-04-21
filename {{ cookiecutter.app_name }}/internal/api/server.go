package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	logger "github.com/sanservices/apilogger/v2"
	"go.uber.org/fx"
	"{{ cookiecutter.module_name }}/settings"
)

// NewServer creates new echo server object and registers start and end of lifecycle of app
// to start echo on start and gracefully shut it down on exit
func NewServer(lc fx.Lifecycle, cfg *settings.Settings) *echo.Echo {
	e := echo.New()

	// avoid any native logging of echo, because we use custom library for logging
	e.HideBanner = true        // don't log the banner on startup
	e.HidePort = true          // hide log about port server started on
	e.Logger.SetLevel(log.OFF) // disable echo#Logger

	// Add hook to start server whenever app starts
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infof(ctx, logger.LogCatRouterInit, "server running on port: %d", cfg.Service.Port)
			go func() {
				err := e.Start(fmt.Sprintf(":%d", cfg.Service.Port))
				if err != nil {
					logger.Fatalf(ctx, logger.LogCatStartUp, "Error while starting server: %s", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info(ctx, logger.LogCatDebug, "Server is shutting down")
			return e.Shutdown(ctx)
		},
	})

	return e
}
