package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler interface {
	RegisterRoutes(e *echo.Group)
}

func New() *echo.Echo {
	e := echo.New()

	e.HideBanner = true        // Don't log the banner on startup
	e.HidePort = true          // Hide log about the port when server i starting up
	e.Logger.SetLevel(log.OFF) // disable echo#Logger

	return e
}
