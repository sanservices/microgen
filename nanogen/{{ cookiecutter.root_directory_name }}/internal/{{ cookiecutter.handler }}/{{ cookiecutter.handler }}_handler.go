package handler

import (
	"function/internal/repository"
	{% if cookiecutter.use_cache == 'y' %}"function/internal/repository/redis"{% endif %}

	ddtracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// init starts the Datadog tracer once when the function package loads. Basic
// configuration (service name, env, version, agent address) is read from the
// standard DD_* environment variables; the tracer flushes on process exit.
func init() {
	ddtracer.Start()
}

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
