package v1

import (
	errs2 "github.com/sanservices/apicore/errs"
	"github.com/sanservices/apicore/helper"
	_ "{{ cookiecutter.module_name }}/internal/api/v1/swagger" // statik file

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
		return helper.RespondError(c, http.StatusInternalServerError, errs2.NewNoTemplateErr())
	}

	indexHTML, err := ioutil.ReadAll(index)
	if err != nil {
		logger.Info(ctx, logger.LogCatFileRead, err)
		return helper.RespondError(c, http.StatusInternalServerError, errs2.NewNoTemplateErr())
	}

	tmpl, err := template.New("index.html").Parse(string(indexHTML))
	if err != nil {
		logger.Info(ctx, logger.LogCatTemplateExec, err)
		return helper.RespondError(c, http.StatusInternalServerError, errs2.NewNoTemplateErr())
	}

	data := DocsModel{
		ServicePrefix: h.cfg.Service.PathPrefix,
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		logger.Info(ctx, logger.LogCatTemplateExec, err)
		return helper.RespondError(c, http.StatusInternalServerError, errs2.NewExecTemplateErr())
	}

	return c.HTML(http.StatusOK, buf.String())
}
