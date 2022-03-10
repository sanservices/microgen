package healthcheck

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/sanservices/apicore/helper"
	logger "github.com/sanservices/apilogger/v2"
	"github.com/sanservices/kit/pprofdebug"
	"net/http"
	"os"
	"time"

	"{{ cookiecutter.module_name }}/goutils/settings"
)

// Handler is the handler for healthchecks
type Handler struct {
	debug bool
}

// healthcheck response
type healthcheck struct {
	Host     string    `json:"host"`
	Datetime time.Time `json:"datetime"`
	Version  string    `json:"version"`
}

// healthcheck response
//
// swagger:response healthcheckRS
type healthCheckResponse struct {
	// in: body
	Body struct {
		// Example: localhost.abf7b06b4ac1
		Host string `json:"host"`
		// Example: 2020-04-01T12:00:00Z
		Datetime time.Time `json:"datetime"`
		// Example: 1.0.0
		Version string `json:"version"`
	}
}

// NewHandler is healthcheck Handler constructor
func NewHandler(config *settings.Settings) *Handler {
	return &Handler{
		debug: config.Service.Debug,
	}
}

// RegisterRoutes initializes api routes
func (h *Handler) RegisterRoutes(e *echo.Group) {
	// healthcheck shows the health and version of the service.
	// swagger:route GET /healthcheck  healthcheck
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
	//  200: healthcheckRS
	e.GET("/healthcheck", h.healthCheck)

	if !h.debug {
		return
	}

	pprofdebug.WrapGroup("", e)
}

func (h *Handler) healthCheck(c echo.Context) error {

	ctx := c.Request().Context()
	sp, _ := opentracing.StartSpanFromContext(ctx, "base-healthCheck")
	defer sp.Finish()

	host, err := os.Hostname()
	if err != nil {
		logger.Error(ctx, logger.LogCatHealth, err)
		return helper.RespondError(c, http.StatusInternalServerError, err)
	}

	resp := healthcheck{
		Host:     host,
		Datetime: time.Now(),
		Version:  "0.0.1",
	}

	return helper.RespondOk(c, resp)
}
