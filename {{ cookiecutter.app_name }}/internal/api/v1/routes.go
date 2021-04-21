package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"
	logger "github.com/sanservices/apilogger/v2"
	_ "{{ cookiecutter.module_name }}/internal/api/v1/swagger" // statik file
	"net/http"
)

// RegisterRoutes initializes api v1 routes
func (h *Handler) RegisterRoutes(e *echo.Group) {
	statikFS, err := fs.New()
	// Try to get swagger from statik
	if err != nil {
		// Log error if it doesn't work
		logger.Error(context.TODO(), logger.LogCatRouterInit, err)
	} else {
		// Add swagger and other static files routes to api
		e.GET("/v1/docs/*", echo.WrapHandler(http.StripPrefix("/v1/docs", http.FileServer(statikFS))))
	}

	// Docs endpoint
	e.GET("/v1/docs", h.getDocs)

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
