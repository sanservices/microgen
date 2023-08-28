package statushandler

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
	log "github.com/sanservices/apilogger/v2"
)

func (h StatusHandler) RegisterRoutes(e *echo.Group) {

	// Healthcheck shows the health and version of the service.
	// Swagger:route GET /healthcheck
	//
	// Shows the health and version of the service.
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  200: healthcheck
	e.GET("/healthcheck", h.healthCheck)
}

func (h *StatusHandler) healthCheck(c echo.Context) error {
	ctx := c.Request()

	host, err := os.Hostname()
	if err != nil {
		log.Error(ctx.Context(), log.LogCatHealth, err)
		return helper.RespondError(c, http.StatusInternalServerError, err)
	}

	out := Healthcheck{
		Host:     host,
		Datetime: time.Now(),
	}

	return helper.RespondOk(c, out)
}
