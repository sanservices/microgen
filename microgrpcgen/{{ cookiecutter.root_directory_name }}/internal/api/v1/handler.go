package v1

import (
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.module_name }}-proto/pb"
	"github.com/sanservices/apicore/validator"
	"{{ cookiecutter.module_name }}/config"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
)

type Handler struct {
	pb.UnimplementedUserServer
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
