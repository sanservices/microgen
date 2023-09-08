package service

import (
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}"
	"{{ cookiecutter.module_name }}/config"
)

type Service struct {
	config *config.Config
	repo {{ cookiecutter.app_name }}.Repository
}

// New returns service instance
func New(config *config.Config, repo {{ cookiecutter.app_name }}.Repository) {{ cookiecutter.app_name }}.Service {
	return &Service{
		config: config,
		repo:   repo,
	}
}
