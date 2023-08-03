package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sanservices/apilogger/v2"
	"{{ cookiecutter.module_name }}/config"
)

var (
	// ErrInvalidDBEngine database engine is not supported
	ErrInvalidDatabaseEngine = errors.New("unsopported database engine")
)

func New(ctx context.Context, config *config.Config) (*sqlx.DB, error) {
	if config.Database.Engine == "mysql" {
		return CreateMySqlConnection(ctx, config)
	} else if config.Database.Engine == "oracle" {
		return CreateOracleConnection(ctx, config)
	}

	return nil, ErrInvalidDatabaseEngine
}

func CreateMySqlConnection(ctx context.Context, config *config.Config) (*sqlx.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	log.Info(ctx, log.LogCatDatastoreConnect, "Connecting to database...")

	db, err := sqlx.ConnectContext(ctx, "mysql", conn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(5 * time.Second))
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(5)

	log.Info(ctx, log.LogCatDatastoreConnect, "Connected to database")
	return db, nil
}

func CreateOracleConnection(ctx context.Context, config *config.Config) (*sqlx.DB, error) {
	return nil, nil
}
