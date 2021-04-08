package v1

import (
	_ "{{cookiecutter.module_name}}/internal/api/v1/swagger" // statik file

	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}"
	"{{cookiecutter.module_name}}/internal/api"
	"{{cookiecutter.module_name}}/settings"

	"context"
	"github.com/rakyll/statik/fs"
	logger "github.com/sanservices/apilogger/v2"
	"net/http"
)

// Handler handles v1 routes
type Handler struct {
	cfg       *settings.Settings
	service   {{cookiecutter.main_domain}}.Service
	validator *api.Validator
	statikFS  http.FileSystem
}

// NewHandler initialize main *Handler
func NewHandler(cfg *settings.Settings, svc {{cookiecutter.main_domain}}.Service, validator *api.Validator) *Handler {
	statikFS, err := fs.New()
	// Try to get swagger from statik
	if err != nil {
		// Log error if it doesn't work
		logger.Error(context.TODO(), logger.LogCatRouterInit, err)
		panic(err)
	}

	return &Handler{
		cfg:       cfg,
		service:   svc,
		validator: validator,
		statikFS:  statikFS,
	}
}
