package api

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxValueKey string

const (
	apiKey        ctxValueKey = "api-key"
	requestIDKey  ctxValueKey = "x-request-id"
	remoteAddrKey ctxValueKey = "remote-address"
)

func handleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.Header.Get(string(requestIDKey))
		key := r.Header.Get(string(apiKey))

		if id == "" {
			id = uuid.New().String()
		}

		addr := r.Header.Get("x-real-ip")
		if addr == "" {
			if r.RemoteAddr != "" {
				addr = r.RemoteAddr
			}
		}

		ctx = context.WithValue(ctx, requestIDKey, id)
		ctx = context.WithValue(ctx, apiKey, key)
		ctx = context.WithValue(ctx, remoteAddrKey, addr)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
