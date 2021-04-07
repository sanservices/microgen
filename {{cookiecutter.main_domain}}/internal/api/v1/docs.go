package v1

import (
	"{{cookiecutter.module_name}}/internal/api"
	"{{cookiecutter.module_name}}/internal/errs"
	"bytes"
	"github.com/labstack/echo/v4"
	logger "github.com/sanservices/apilogger/v2"
	"net/http"
	"text/template"
)

// DocsModel is handler struct for docs
type DocsModel struct {
	ServicePrefix string
}

func (h *handler) getDocs(c echo.Context) error {
	ctx := c.Request().Context()

	tmpl, err := template.New("index.html").ParseFiles("static/swaggerui/v1/index.html")
	if err != nil {
		logger.Info(ctx, logger.LogCatStartUp, err)
		return api.RespondError(c, http.StatusInternalServerError, errs.NewNoTemplateErr())
	}

	data := DocsModel{
		ServicePrefix: h.cfg.Service.PathPrefix,
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		logger.Info(ctx, logger.LogCatStartUp, err)
		return api.RespondError(c, http.StatusInternalServerError, errs.NewExecTemplateErr())
	}
	return c.HTML(http.StatusOK, buf.String())
}
