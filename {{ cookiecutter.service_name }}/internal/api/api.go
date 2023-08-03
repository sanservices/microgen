package api

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(e *echo.Group)
}

func New() *echo.Echo {
	e := echo.New()

	e.HideBanner = true // Don't log the banner on startup
	e.HidePort = true   // Hide log about the port when server i starting up

	return e
}

func RegisterRoutes(e *echo.Echo, handlers []Handler) {
	for _, h := range handlers {
		h.RegisterRoutes(e.Group(""))
	}

	for _, r := range e.Routes() {
		log.Printf("[%s] %s", r.Method, r.Path)
	}
}
