package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
)

func (h *Handler) getThing(c echo.Context) error {
	queryThingID := c.QueryParam("id")
	thingID, err := strconv.Atoi(queryThingID)
	if err != nil || thingID < 0 {
		return helper.RespondError(c, http.StatusBadRequest, errors.New("id query param must be unsigned integer"))
	}

	thing, err := h.service.GetUser(c.Request().Context(), uint(thingID))
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, errors.New("something went wrong"))
	}
	return helper.RespondOk(c, thing)
}

func (h *Handler) createThing(c echo.Context) error {
	// Create a new thing
	return helper.RespondOk(c, "OK")
}
