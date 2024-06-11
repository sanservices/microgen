package v1

import (
	"github.com/sanservices/apicore/validator"
	"{{ cookiecutter.module_name }}/config"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
)

type Handler struct {
	config   *config.Settings
	service  {{ cookiecutter.service_name }}.Service
	validate *validator.Validator
}

func New(config *config.Settings, service {{ cookiecutter.service_name }}.Service) *Handler {
	return &Handler{
		config:   config,
		service:  service,
		validate: validator.NewValidator(),
	}
}
