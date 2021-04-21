package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"
	apicoreMW "github.com/sanservices/apicore/middleware"
	logger "github.com/sanservices/apilogger/v2"
	"{{ cookiecutter.module_name }}/settings"
)

// RegisterRoutes iterates over handlers and registers them in given echo server instance
func RegisterRoutes(cfg *settings.Settings, e *echo.Echo, handlers []Handler) {
	e.Use(apicoreMW.SetCustomHeaders)
	e.Use(apicoreMW.EnrichContext)
	e.Use(apicoreMW.RequestLogger)
	e.Use(echoMW.Recover())

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
