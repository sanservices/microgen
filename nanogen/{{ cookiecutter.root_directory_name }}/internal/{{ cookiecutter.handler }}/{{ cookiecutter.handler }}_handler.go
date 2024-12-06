package handler

import (
	"function/internal/repository"
	{% if cookiecutter.use_cache == 'y' %}"function/internal/repository/redis"{% endif %}
)

type Handler struct {
	repo repository.Repository
	{% if cookiecutter.use_cache == 'y' %} redis redis.Cache {% endif %}
}

func New(repo repository.Repository {% if cookiecutter.use_cache == 'y' %}, redis redis.Cache {% endif %}) *Handler {
	return &Handler{
		repo: repo,
		{% if cookiecutter.use_cache == 'y' %} redis: redis, {% endif %}
	}
}
