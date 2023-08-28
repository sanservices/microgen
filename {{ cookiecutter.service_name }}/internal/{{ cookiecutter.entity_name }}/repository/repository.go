package repository

import (
	"context"
	
	{% if cookiecutter.use_cache  == 'y' %}"github.com/go-redis/redis/v8"{% endif %}
	{% if cookiecutter.use_database == 'y' %}"github.com/jmoiron/sqlx"{% endif %}
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.entity_name }}"
)

type Repository struct {
	{% if cookiecutter.use_database == 'y' %}db    *sqlx.DB{% endif %}
	{% if cookiecutter.use_cache == 'y' %}cache *redis.Client{% endif %}
}

// New repository instance
func New(ctx context.Context, {% if cookiecutter.use_database == 'y' %}db *sqlx.DB,{% endif %} {% if cookiecutter.use_cache == 'y' %}cache *redis.Client{% endif %}) {{ cookiecutter.entity_name }}.Repository {
	return &Repository{
		{% if cookiecutter.use_database == 'y' %}db:    db,{% endif %}
		{% if cookiecutter.use_cache == 'y' %}cache: cache,{% endif %}
	}
}
