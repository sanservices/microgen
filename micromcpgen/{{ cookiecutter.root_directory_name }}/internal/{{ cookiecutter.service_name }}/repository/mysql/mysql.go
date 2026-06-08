package mysql

import (
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
)

// Mysql connection
type Mysql struct {
	db *sqlx.DB
}

// New returns an instance of Mysql connection
func New(db *sqlx.DB) (m *Mysql) {
	return &Mysql{
		db: db,
	}
}
