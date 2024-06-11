package db

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sanservices/kit/database"
	"{{ cookiecutter.module_name }}/config"
)

var (
	ErrInvalidDatabaseEngine = errors.New("unsopported database engine")
)

func New(ctx context.Context, config *config.Settings) (*sqlx.DB, error) {
	switch config.Database.Engine {
	case "mysql":
		return database.CreateMySqlConnection(ctx, config.Database)

	case "oracle":
		return database.CreateOracleConnection(ctx, config.Database)

	case "sqlite":
		return database.CreateSqliteConnection(ctx, config.Database)

	default:
		return nil, ErrInvalidDatabaseEngine
	}
}
