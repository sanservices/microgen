package api

import (
	"{{cookiecutter.module_name}}/settings"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logger "github.com/sanservices/apilogger/v2"
)

// RegisterRoutes iterates over handlers and registers them in given echo server instance
func RegisterRoutes(cfg *settings.Settings, e *echo.Echo, handlers []Handler) {
	e.Use(enrichContext)
	e.Use(requestLogger)
	e.Use(middleware.Recover())

	for _, h := range handlers {
		h.RegisterRoutes(e.Group(""))
	}

	if cfg.Service.Debug {
		var routeList []string

		// Make an array of routes(formatted strings)
		for _, r := range e.Routes() {
			routeList = append(routeList, fmt.Sprintf(" [%s] %s ;", r.Method, r.Path))
		}
		logger.InfoWF(context.TODO(), logger.LogCatRouterInit, &logger.Fields{"routes": routeList})
	}
}
