package service

import (
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.entity_name }}"
)

type Service struct {
	repo {{ cookiecutter.entity_name }}.Repository
}

// New returns service instance
func New(repo {{ cookiecutter.entity_name }}.Repository) {{ cookiecutter.entity_name }}.Service {
	return &Service{
		repo: repo,
	}
}
