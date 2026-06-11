package v1

import "github.com/labstack/echo/v4"

// RegisterRoutes wires the swagger docs endpoints onto the echo server, letting
// the v1 Handler satisfy the api.Handler interface. The REST API itself is served
// by the gRPC-gateway catch-all mounted in main.go, so only the docs routes are
// registered here.
func (h Handler) RegisterRoutes(e *echo.Group) {
	e.GET("/v1/docs", h.getSwaggerIndex)
	e.GET("/v1/docs/swagger.yml", h.getSwaggerSchema)
}
