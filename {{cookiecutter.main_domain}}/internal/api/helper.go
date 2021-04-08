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
	var response errorResponse
	if err, ok := respErr.(errs.ServiceError); ok {
		response.Errors = append(response.Errors, err)
	} else {
		switch httpStatusCode {
		case http.StatusBadRequest:
			response.Errors = append(response.Errors, errs.NewInputMalformedErr(respErr))
		default:
			response.Errors = append(response.Errors, errs.NewInternalErr(respErr))
		}
	}

	return c.JSON(httpStatusCode, response)
}

// RespondOk sends a json success response
func RespondOk(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, &successResponse{Data: data})
}
