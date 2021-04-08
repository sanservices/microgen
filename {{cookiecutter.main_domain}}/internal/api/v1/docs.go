package v1

import (
	_ "{{cookiecutter.module_name}}/internal/api/v1/swagger" // statik file


	"{{cookiecutter.module_name}}/internal/api"
	"{{cookiecutter.module_name}}/internal/errs"
	"bytes"
	"github.com/labstack/echo/v4"
	logger "github.com/sanservices/apilogger/v2"
	"io/ioutil"
	"net/http"
	"text/template"
)

// DocsModel is handler struct for docs
type DocsModel struct {
	ServicePrefix string
}

func (h *Handler) getDocs(c echo.Context) error {
	ctx := c.Request().Context()

	index, err := h.statikFS.Open("/index.html")
	if err != nil {
		logger.Info(ctx, logger.LogCatFileRead, err)
		return api.RespondError(c, http.StatusInternalServerError, errs.NewNoTemplateErr())
	}

	indexHTML, err := ioutil.ReadAll(index)
	if err != nil {
		logger.Info(ctx, logger.LogCatFileRead, err)
		return api.RespondError(c, http.StatusInternalServerError, errs.NewNoTemplateErr())
	}

	tmpl, err := template.New("index.html").Parse(string(indexHTML))
	if err != nil {
		logger.Info(ctx, logger.LogCatTemplateExec, err)
		return api.RespondError(c, http.StatusInternalServerError, errs.NewNoTemplateErr())
	}

	data := DocsModel{
		ServicePrefix: h.cfg.Service.PathPrefix,
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		logger.Info(ctx, logger.LogCatTemplateExec, err)
		return api.RespondError(c, http.StatusInternalServerError, errs.NewExecTemplateErr())
	}

	return c.HTML(http.StatusOK, buf.String())
}
