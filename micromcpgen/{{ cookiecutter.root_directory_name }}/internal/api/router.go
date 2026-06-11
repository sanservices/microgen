package api

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"
	ddecho "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	apicoreMW "github.com/sanservices/apicore/middleware"
	logger "github.com/sanservices/apilogger/v2"
)

// RegisterRoutes iterates over handlers and registers them in given echo server instance
func RegisterRoutes(e *echo.Echo, handlers []Handler) {
	// Datadog APM tracing for all incoming HTTP requests
	e.Use(ddecho.Middleware())

	e.Use(apicoreMW.SetCustomHeaders)
	e.Use(apicoreMW.EnrichContext)
	e.Use(apicoreMW.RequestLogger)
	e.Use(echoMW.Recover())

	for _, h := range handlers {
		h.RegisterRoutes(e.Group(""))
	}

	var routeList []string

	// Make an array of routes(formatted strings)
	for _, r := range e.Routes() {
		routeList = append(routeList, fmt.Sprintf(" [%s] %s ;", r.Method, r.Path))
	}
	logger.InfoWF(context.TODO(), logger.LogCatRouterInit, &logger.Fields{"routes": routeList})
}
