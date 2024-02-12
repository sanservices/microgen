package repository

import (
	"context"
	
	"errors"

	{% if cookiecutter.use_database == 'y' %}"github.com/jmoiron/sqlx"{% endif %}
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
	"{{ cookiecutter.module_name }}/config"

	{% if cookiecutter.use_database != 'n' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository/mysql"{% endif %}
	{% if cookiecutter.use_database != 'n' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository/oracle"{% endif %}
	{% if cookiecutter.use_database != 'n' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/repository/sqlite"{% endif %}
)

// ErrInvalidDBEngine database engine is not supported
var ErrInvalidDBEngine = errors.New("unsoported or missing database engine")

// New repository instance
func New(ctx context.Context, cfg *config.Settings{% if cookiecutter.use_database == 'y' %} ,db *sqlx.DB,{% endif %}) ({{ cookiecutter.service_name }}.Repository, error) {

	{% if cookiecutter.use_database == 'y' %}
	switch cfg.Database.Engine {
		case "mysql":
			return mysql.New(db), nil
		case "oracle":
			return oracle.New(db), nil
		case "sqlite":
			repo := sqlite.New(db)
			return repo, repo.PopulateSchema(ctx)
		default:
			return nil, ErrInvalidDBEngine
	}
	{% else %}
	return nil, nil
	{% endif %}
}
