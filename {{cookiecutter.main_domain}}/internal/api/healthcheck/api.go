package healthcheck

import (
	"{{cookiecutter.module_name}}/internal/api"
	"github.com/labstack/echo/v4"
	logger "github.com/sanservices/apilogger/v2"
	"net/http"
	"os"
	"time"
)

// handler is the handler for healthchecks
type handler struct {
}

// NewHandler is healthcheck handler constructor
func NewHandler() *handler {
	return &handler{}
}

// RegisterRoutes initializes api routes
func (h *handler) RegisterRoutes(e *echo.Group) {
	e.GET("/healthcheck", h.healthCheck)
}

func (h *handler) healthCheck(c echo.Context) error {
	resp := struct {
		Host     string    `json:"host"`
		Datetime time.Time `json:"datetime"`
	}{
		Datetime: time.Now(),
	}

	var err error
	resp.Host, err = os.Hostname()
	if err != nil {
		logger.Error(c.Request().Context(), logger.LogCatHealth, err)
		return api.RespondError(c, http.StatusOK, err)
	}

	return api.RespondOk(c, resp)
}
