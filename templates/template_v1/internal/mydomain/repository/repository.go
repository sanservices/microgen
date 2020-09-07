package repository

import (
	"context"
	"errors"
	"goproposal/internal/mydomain"
	"goproposal/internal/mydomain/repository/mocked"

	"github.com/san-services/apicore/apisettings"
	"github.com/san-services/apilogger"
)

// New constructs the repository
func New(ctx context.Context, dbConfig apisettings.Database) (mydomain.Repository, error) {
	lg := apilogger.New(ctx, "")

	switch dbConfig.Engine {
	case "mocked":
		return mocked.New(ctx, dbConfig)
	// case "mysql":
	// 	return mysql.New(ctx, dbConfig), nil
	default:
		err := errors.New("Unsupported or missing database engine")
		lg.Error(apilogger.LogCatReadConfig, err)
		return nil, err
	}
}
