package v1

import "github.com/labstack/echo/v4"

// Within the RegisterRoutes method, you can use the provided Echo Group instance (e) to define your routes
// using HTTP methods like GET, POST, PUT, DELETE, etc., along with corresponding route paths and handler functions.
// For example:
//
//	e.GET("/entity", h.GetEntityHandler)
//	e.POST("/entity", h.CreateEntityHandler)
//	e.PUT("/entity/:id", h.UpdateEntityHandler)
//	e.DELETE("/entity/:id", h.DeleteEntityHandler)
func (h Handler) RegisterRoutes(e *echo.Group) {

	// Swagger routes
	e.GET("/v1/docs", h.getSwaggerIndex)
	e.GET("/v1/docs/swagger.yml", h.getSwaggerSchema)

	// routes for version 1: /v1/...
	v1 := e.Group("/v1")

	v1.GET("/entity", h.getThing)
	v1.POST("/entity", h.createThing)
}
