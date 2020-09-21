// Goproposal
//
// Base/example service to help with efficiently
// and excellently producing of new microservices
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
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
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/san-services/apicore"
	"github.com/san-services/apicore/apisettings"
	"github.com/san-services/apilogger"
)

func main() {
	ctx := context.Background()
	lg := apilogger.New(ctx, "")
	wait := time.Second * 15

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

	srv := &http.Server{
		Addr:         "0.0.0.0" + portToListen,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Router, // Pass our instance of gorilla/mux in.
	}

	go func() {
		log.Printf("server started at http://localhost:%d", config.Service.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctxWait, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctxWait)
	log.Println("shutting down")
	os.Exit(0)
}
