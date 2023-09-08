package oracle

import (
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"

	{% if cookiecutter.use_cache == 'y' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}"{% endif %}
)

// Mysql connection
type Oracle struct {
	db *sqlx.DB
	{% if cookiecutter.use_cache == 'y' %}cache {{ cookiecutter.app_name }}.Cache{% endif %}
}

// New returns an instance of Mysql connection
func New(db *sqlx.DB {% if cookiecutter.use_cache == 'y' %}, cache {{ cookiecutter.app_name }}.Cache{% endif %}) (o *Oracle) {
	return &Oracle{
		db: db,
		{% if cookiecutter.use_cache == 'y' %}cache: cache,{% endif %}
	}
}
