package v1

import (
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}"
	"{{cookiecutter.module_name}}/internal/api"
	"{{cookiecutter.module_name}}/settings"
)

// handler handles v1 routes
type handler struct {
	cfg       *settings.Settings
	service   {{cookiecutter.main_domain}}.Service
	validator *api.Validator
}

// NewHandler initialize main *handler
func NewHandler(cfg *settings.Settings, svc {{cookiecutter.main_domain}}.Service, validator *api.Validator) *handler {
	return &handler{
		cfg:       cfg,
		service:   svc,
		validator: validator,
	}
}
