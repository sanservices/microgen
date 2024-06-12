package repository

import (
	"context"
	"errors"
	"function/internal/config"
	"function/internal/repository/mysql"
	"function/internal/repository/oracle"

	"github.com/jmoiron/sqlx"
)

// Repository is the contract interface for the repository layer,
// responsible for data storage and retrieval.
type Repository interface{}

func New(ctx context.Context, cfg *config.Settings {% if cookiecutter.use_database == 'y' %}, db *sqlx.DB{% endif %}) (Repository, error) {

	{% if cookiecutter.use_database == 'y' %}
	switch cfg.Database.Engine {
	case "mysql":
		return mysql.New(db), nil
	case "oracle":
		return oracle.New(db), nil
	default:
		return nil, errors.ErrUnsupported
	}
	{% endif %}
}
