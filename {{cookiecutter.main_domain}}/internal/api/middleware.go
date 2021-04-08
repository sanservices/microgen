package api

import (
	"context"
	"github.com/dchest/uniuri"
	"github.com/labstack/echo/v4"
	logger "github.com/sanservices/apilogger/v2"
	"net/http"
	"time"
)

// Sets some custom headers
func setCustomHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := next(c)

		id, ok := c.Request().Context().Value("x-request-id").(string)
		if !ok {
			id = ""
		}

		c.Response().Header().Set("X-Request-Id", id)

		return res
	}
}

// Enriches the context and sets context to request
func enrichContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		id := r.Header.Get(logger.RequestIDKey)
		key := r.Header.Get(logger.APIKEY)
		session := r.Header.Get(logger.SessionIDKey)

		if id == "" {
			id = uniuri.New()
		}

		addr := r.Header.Get("x-real-ip")
		if addr == "" {
			if r.RemoteAddr != "" {
				addr = r.RemoteAddr
			}
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, logger.APIKEY, key)
		ctx = context.WithValue(ctx, logger.RequestIDKey, id)
		ctx = context.WithValue(ctx, logger.RemoteAddrKey, addr)
		ctx = context.WithValue(ctx, logger.SessionIDKey, session)
		ctx = context.WithValue(ctx, logger.StartTime, time.Now())

		c.SetRequest(r.WithContext(ctx)) // set context to request

		return next(c)
	}
}

// Logs requests
func requestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		// Call next handler
		err := next(c)

		logger.InfoWF(r.Context(), logger.LogCatReqPath, &logger.Fields{
			"status": http.StatusOK,
			"url":    r.URL,
			"method": r.Method,
		})

		return err
	}
}
