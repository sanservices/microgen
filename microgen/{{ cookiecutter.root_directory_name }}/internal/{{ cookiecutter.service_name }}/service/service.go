package service

import (
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
	"{{ cookiecutter.module_name }}/config"
)

type Service struct {
	config *config.Settings
	repo {{ cookiecutter.service_name }}.Repository
	{% if cookiecutter.use_cache != 'n' %}redis {{ cookiecutter.service_name }}.Cache{% endif %}
}

// New returns service instance
func New(config *config.Settings, repo {{ cookiecutter.service_name }}.Repository {% if cookiecutter.use_cache != 'n' %}, redis {{ cookiecutter.service_name }}.Cache{% endif %}) {{ cookiecutter.service_name }}.Service {
	return &Service{
		config: config,
		repo:   repo,
		{% if cookiecutter.use_cache != 'n' %}redis: redis,{% endif %}
	}
}
