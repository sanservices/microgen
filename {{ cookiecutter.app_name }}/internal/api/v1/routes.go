package v1

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes initializes api v1 routes
func (h *Handler) RegisterRoutes(e *echo.Group) {

	// Swagger routes
	e.GET("/v1/docs", h.getSwaggerIndex)
	e.GET("/v1/docs/swagger.yml", h.getSwaggerSchema)

	// swagger:route GET /thing things getThingRQ
	//
	// Retrieves a thing
	//
	// Retrieves a thing
	//
	// responses:
	//    200: getThingRS
	//    400: badRequestRS
	//	  500: serverErrorRS
	e.GET("/v1/thing", h.getThing)
}
