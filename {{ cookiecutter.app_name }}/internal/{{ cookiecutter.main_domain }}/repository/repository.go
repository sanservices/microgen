package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	logger "github.com/sanservices/apilogger/v2"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository/fixture"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository/mock"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository/mysql"
	"{{ cookiecutter.module_name }}/settings"
	"time"
)

// New constructs the repository
func New(ctx context.Context, cfg *settings.Settings) (repo {{ cookiecutter.main_domain }}.Repository) {
	var err error
	defer func() {
		if err != nil {
			logger.Fatal(ctx, logger.LogCatStartUp, "could not create repository", err)
		}
	}()

	switch cfg.DB.Engine {
	case "mock":
		return mock.NewWithExpectations()
	case "fixture":
		return fixture.New()
	case "mysql":
		connectionString := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Name,
		)

		mysqlConn, err := sqlx.Open("mysql", connectionString)
		if err != nil {
			logger.Error(ctx, logger.LogCatDatastoreConnect, err)
			return nil
		}

		mysqlConn.SetMaxOpenConns(25) // Set some limits for connection pool
		mysqlConn.SetMaxIdleConns(15)
		mysqlConn.SetConnMaxLifetime(5 * time.Minute)

		err = mysqlConn.Ping()
		if err != nil {
			logger.Error(ctx, logger.LogCatDatastoreConnect, err)
			return nil
		}

		return mysql.New(mysqlConn)
	default:
		err := errors.New("unsupported or missing database engine")
		logger.Error(ctx, logger.LogCatReadConfig, err)
		return nil
	}
}
