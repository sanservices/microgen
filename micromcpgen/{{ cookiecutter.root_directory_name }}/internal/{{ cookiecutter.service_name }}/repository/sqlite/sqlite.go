package sqlite

import (
	"context"
	_ "embed"
	"github.com/jmoiron/sqlx"
)

type Sqlite struct {
	db *sqlx.DB
}

//go:embed schema.sql
var createSchemaStmt string

func New(db *sqlx.DB) *Sqlite {
	return &Sqlite{
		db: db,
	}
}

func (sl Sqlite) PopulateSchema(ctx context.Context) error {
	_, err := sl.db.ExecContext(ctx, createSchemaStmt)
	return err
}
