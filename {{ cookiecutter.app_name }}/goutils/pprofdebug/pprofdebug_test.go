package pprofdebug

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestWrapGroup(t *testing.T) {
	routes := []struct {
		method string
		path   string
	}{
		{method: "GET", path: "/debug/pprof/"},
		{method: "GET", path: "/debug/pprof/heap"},
		{method: "GET", path: "/debug/pprof/goroutine"},
		{method: "GET", path: "/debug/pprof/block"},
		{method: "GET", path: "/debug/pprof/threadcreate"},
		{method: "GET", path: "/debug/pprof/cmdline"},
		{method: "GET", path: "/debug/pprof/profile"},
		{method: "GET", path: "/debug/pprof/trace"},
		{method: "GET", path: "/debug/pprof/mutex"},
		{method: "GET", path: "/debug/pprof/symbol"},
		{method: "POST", path: "/debug/pprof/symbol"},
	}
	e := echo.New()

	WrapGroup("", e.Group(""))

	for _, route := range routes {
		found := false

		for _, r := range e.Routes() {
			if route.method == r.Method && route.path == r.Path {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Incomplete list routes, missing: %v", route)
		}
	}
}
