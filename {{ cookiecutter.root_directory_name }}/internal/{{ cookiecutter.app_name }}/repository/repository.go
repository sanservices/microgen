package repository

import (
	"context"
	
	"errors"

	{% if cookiecutter.use_database == 'y' %}"github.com/jmoiron/sqlx"{% endif %}
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}"
	"{{ cookiecutter.module_name }}/config"

	{% if cookiecutter.use_database != 'n' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/repository/mysql"{% endif %}
	{% if cookiecutter.use_database != 'n' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/repository/oracle"{% endif %}
	{% if cookiecutter.use_database != 'n' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/repository/sqlite"{% endif %}

	{% if cookiecutter.use_cache  == 'y' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}/repository/redis" {% endif %}
)

// ErrInvalidDBEngine database engine is not supported
var ErrInvalidDBEngine = errors.New("unsoported or missing database engine")

// New repository instance
func New(ctx context.Context, cfg *config.Config{% if cookiecutter.use_database == 'y' %} ,db *sqlx.DB,{% endif %}) ({{ cookiecutter.app_name }}.Repository, error) {
	
	{% if cookiecutter.use_cache == 'y' %}
	cache, err := redis.New(ctx, cfg)
	if err != nil {
		return nil, err
	}
	{% endif %} 

	{% if cookiecutter.use_database == 'y' %}
	switch cfg.Database.Engine {
		case "mysql":
			return mysql.New(db {% if cookiecutter.use_cache == 'y' %}, cache{% endif %}), nil
		case "oracle":
			return oracle.New(db{% if cookiecutter.use_cache == 'y' %}, cache{% endif %}), nil
		case "sqlite":
			repo := sqlite.New(db{% if cookiecutter.use_cache == 'y' %}, cache{% endif %})
			return repo, repo.PopulateSchema(ctx)
		default:
			return nil, ErrInvalidDBEngine
	}
	{% else %}
	return nil, nil
	{% endif %}
}
