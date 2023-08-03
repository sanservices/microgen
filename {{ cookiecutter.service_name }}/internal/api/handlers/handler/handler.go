package handler

import (
	"github.com/sanservices/apicore/validator"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.entity_name }}"
)

type Handler struct {
	service  {{ cookiecutter.entity_name }}.Service
	validate *validator.Validator
}

func New(service {{ cookiecutter.entity_name }}.Service) *Handler {
	return &Handler{
		service:  service,
		validate: validator.NewValidator(),
	}
}
