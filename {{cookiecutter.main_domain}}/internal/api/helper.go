package api

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"{{cookiecutter.module_name}}/internal/errs"
)

// httpResponse standard json response
type successResponse struct {
	Data interface{} `json:"data"`
}

// JSON response with error
type errorResponse struct {
	Errors []errs.ServiceError `json:"errors"`
}

// RespondError sends a json error response
func RespondError(c echo.Context, httpStatusCode int, respErr error) error {
	ctx := c.Request().Context()

	id, ok := ctx.Value("x-request-id").(string)
	if !ok {
		id = ""
	}

	var response errorResponse
	if err, ok := respErr.(errs.ServiceError); ok {
		response.Errors = append(response.Errors, err)
	} else {
		response.Errors = append(response.Errors, errs.NewInternalErr(respErr))
	}

	c.Response().Header().Set("X-Request-Id", id)

	return c.JSON(httpStatusCode, response)
}

// RespondOk sends a json success response
func RespondOk(c echo.Context, data interface{}) error {
	ctx := c.Request().Context()

	id, ok := ctx.Value("x-request-id").(string)
	if !ok {
		id = ""
	}

	response := &successResponse{Data: data}

	c.Response().Header().Set("X-Request-Id", id)

	return c.JSON(http.StatusOK, response)
}
