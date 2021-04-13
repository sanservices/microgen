package api

import (
	"github.com/labstack/echo/v4"
)

// Handler is an interface that
// each api version's handler should implement
type Handler interface {
	RegisterRoutes(echo *echo.Group)
}
