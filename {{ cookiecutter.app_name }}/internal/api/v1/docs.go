package v1

import (
	"github.com/sanservices/apicore/errs"
	"github.com/sanservices/apicore/helper"

	_ "embed"

	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apilogger/v2"
	"net/http"
	"text/template"
)

//go:embed swagger/index.html
var swaggerIndex string

//go:embed swagger/swagger.yml
var swaggerFile string

// DocsModel is handler struct for docs
type DocsModel struct {
	ServicePrefix string
}

func (h Handler) getSwaggerIndex(c echo.Context) error {
	ctx := c.Request().Context()

	tmpl, err := template.New("index.html").Parse(swaggerIndex)
	if err != nil {
		apilogger.Info(ctx, apilogger.LogCatTemplateExec, err)
		return helper.RespondError(c, http.StatusInternalServerError, errs.NewNoTemplateErr())
	}

	data := DocsModel{}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		apilogger.Info(ctx, apilogger.LogCatTemplateExec, err)
		return helper.RespondError(c, http.StatusInternalServerError, errs.NewExecTemplateErr())
	}

	return c.HTML(http.StatusOK, buf.String())
}

func (h Handler) getSwaggerSchema(c echo.Context) error {
	return c.String(http.StatusOK, swaggerFile)
}
