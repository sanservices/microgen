package handler

import (
	"github.com/sanservices/apicore/validator"
	"{{ cookiecutter.module_name }}/config"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.entity_name }}"
)

type Handler struct {
	config   *config.Config
	service  {{ cookiecutter.entity_name }}.Service
	validate *validator.Validator
}

func New(config *config.Config, service {{ cookiecutter.entity_name }}.Service) *Handler {
	return &Handler{
		config:   config,
		service:  service,
		validate: validator.NewValidator(),
	}
}
