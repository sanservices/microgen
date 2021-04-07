package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"{{cookiecutter.module_name}}/internal/api"
	"net/http"
	"strconv"
)

func (h *handler) getThing(c echo.Context) error {
	queryThingID := c.QueryParam("id")
	thingID, err := strconv.Atoi(queryThingID)
	if err != nil || thingID < 0 {
		return api.RespondError(c, http.StatusBadRequest, errors.New("id query param must be unsigned integer"))
	}

	thing, err := h.service.GetThing(c.Request().Context(), uint(thingID))
	if err != nil {
		return api.RespondError(c, http.StatusInternalServerError, errors.New("something went wrong"))
	}
	return api.RespondOk(c, thing)
}
