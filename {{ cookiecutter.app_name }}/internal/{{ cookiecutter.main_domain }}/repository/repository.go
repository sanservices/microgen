package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"{{ cookiecutter.module_name }}/goutils/settings"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository/mysql"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository/sqlite"
	logger "github.com/sanservices/apilogger/v2"
)

// New constructs the repository
func New(ctx context.Context, cfg *settings.Settings, db *sqlx.DB) ({{ cookiecutter.main_domain }}.Repository, error) {

	switch cfg.DB.Engine {
	case "mysql":
		return mysql.New(db), nil

	case "sqlite":
		repo :=sqlite.New(db)
		return repo, repo.PopulateSchema(ctx)

	default:
		err := errors.New("unsupported or missing database engine")
		logger.Error(ctx, logger.LogCatReadConfig, err)
		return nil, err
	}
}
