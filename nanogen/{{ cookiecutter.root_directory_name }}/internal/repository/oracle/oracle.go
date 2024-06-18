package oracle

import (
	"github.com/jmoiron/sqlx"
)

// Mysql connection
type Oracle struct {
	db *sqlx.DB
}

// New returns an instance of Mysql connection
func New(db *sqlx.DB) (o *Oracle) {
	return &Oracle{
		db: db,
	}
}