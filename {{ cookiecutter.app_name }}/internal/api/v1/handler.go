package v1

import (
	validator2 "github.com/sanservices/apicore/validator"

	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}"
	"{{ cookiecutter.module_name }}/goutils/settings"
)

// Handler handles v1 routes
type Handler struct {
	cfg       *settings.Settings
	service   {{ cookiecutter.main_domain }}.Service
	validator *validator2.Validator
}

// NewHandler initialize main *Handler
func NewHandler(cfg *settings.Settings, svc {{ cookiecutter.main_domain }}.Service, validator *validator2.Validator) *Handler {

	return &Handler{
		cfg:       cfg,
		service:   svc,
		validator: validator,
	}
}
