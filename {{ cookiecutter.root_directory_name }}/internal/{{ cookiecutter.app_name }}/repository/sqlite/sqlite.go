package sqlite

import (
	"context"
	_ "embed"
	"github.com/jmoiron/sqlx"

	{% if cookiecutter.use_cache == 'y' %}"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.app_name }}"{% endif %}
)

type Sqlite struct {
	db *sqlx.DB
	{% if cookiecutter.use_cache == 'y' %}cache {{ cookiecutter.app_name }}.Cache{% endif %}
}

//go:embed schema.sql
var createSchemaStmt string

func New(db *sqlx.DB {% if cookiecutter.use_cache == 'y' %}, cache {{ cookiecutter.app_name }}.Cache{% endif %}) *Sqlite {
	return &Sqlite{
		db: db,
		{% if cookiecutter.use_cache == 'y' %}cache: cache,{% endif %}
	}
}

func (sl Sqlite) PopulateSchema(ctx context.Context) error {
	_, err := sl.db.ExecContext(ctx, createSchemaStmt)
	return err
}
