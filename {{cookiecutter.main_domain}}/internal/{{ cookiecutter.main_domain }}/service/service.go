package service

import (
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}"
)

// Service is a struct able to access all data required
// to perform business logic functions
type Service struct {
	repo {{cookiecutter.main_domain}}.Repository
}

// New constructs and returns a Service struct
func New(repo {{cookiecutter.main_domain}}.Repository) {{cookiecutter.main_domain}}.Service {
	return &Service{
		repo: repo,
	}
}
