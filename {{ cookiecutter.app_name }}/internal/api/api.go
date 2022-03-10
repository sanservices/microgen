package api

import (
	"log"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"
	"github.com/sanservices/apicore/middleware"
	"{{ cookiecutter.module_name }}/goutils/settings"
)

// Handler is an interface that
// each api version's handler should implement
type Handler interface {
	RegisterRoutes(echo *echo.Group)
}

func RegisterRoutes(config *settings.Settings, e *echo.Echo, handlers []Handler) {

	// Enable tracing middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.Use(middleware.SetCustomHeaders)
	e.Use((middleware.EnrichContext))
	e.Use(middleware.RequestLogger)
	e.Use(echoMW.Recover())

	for _, h := range handlers {
		h.RegisterRoutes(e.Group(""))
	}

	for _, r := range e.Routes() {
		log.Printf("[%s] %s", r.Method, r.Path)
	}
}
