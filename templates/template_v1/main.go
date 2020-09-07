// Goproposal
//
// Base/example service to help with efficiently
// and excellently producing of new microservices
//
// Schemes: http
// Host: localhost:8080
// BasePath: /v1
// Version: 1.0
//
// Security:
//     - api_key:
//
// SecurityDefinitions:
//  api_key:
//   type: apiKey
//   name: api-key
//   in: header
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	v1 "goproposal/internal/api/v1"
	"goproposal/internal/mydomain/repository"
	"goproposal/internal/mydomain/service"
	"log"
	"net/http"
	"strconv"

	"github.com/san-services/apicore"
	"github.com/san-services/apicore/apisettings"
	"github.com/san-services/apilogger"
)

func main() {
	ctx := context.Background()
	lg := apilogger.New(ctx, "")

	//Get project configurations.
	config, err := apisettings.Get(ctx, "settings.yml")
	if err != nil {
		lg.Fatal(apilogger.LogCatReadConfig, err)
	}

	// Repository
	repo, err := repository.New(ctx, config.DB)
	if err != nil {
		lg.Fatal(apilogger.LogCatRepoInit, err)
	}

	// Initialize business logic layer.
	service, err := service.New(config.Cache, repo)
	if err != nil {
		lg.Fatal(apilogger.LogCatServiceOutput, err)
	}

	// handlers
	handler := v1.NewHandler(config, service)

	// api
	api := apicore.New(ctx, &config.Service, []apicore.Handler{handler})

	//Start server.
	portToListen := ":" + strconv.Itoa(config.Service.Port)

	lg.Infof(apilogger.LogCatRouterInit, "server running on port: %d", config.Service.Port)
	log.Printf("server started at http://localhost:%d", config.Service.Port)
	log.Fatal(
		http.ListenAndServe(portToListen, api.Router),
	)
}
