package dbfactory

import (
	"context"
	"errors"
	"fmt"
	"log"

	"{{ cookiecutter.module_name }}/goutils/settings"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// ErrInvalidDBEngine database engine is not supported
	ErrInvalidDBEngine = errors.New("unsoported or missing database engine")
)

func New(config *settings.Settings) (*sqlx.DB, error) {

	switch config.DB.Engine {
	case "mysql":
		return CreateMySqlConnection(config)

	case "sqlite":
		return CreateSqliteConnection(config)

	default:
		return nil, ErrInvalidDBEngine
	}
}

func CreateMySqlConnection(config *settings.Settings) (*sqlx.DB, error) {

	var connectionString string
	var db *sqlx.DB
	var err error
	dbConfig := config.DB

	connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	log.Println("Connecting to database...")
	db, err = sqlx.Connect("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}

func CreateSqliteConnection(config *settings.Settings) (*sqlx.DB, error) {
	log.Println("Connecting to database...")
	source := fmt.Sprintf("./%s.db", config.DB.Name)

	db, err := sqlx.Connect("sqlite3", source)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}

func CreateRedisConnection(ctx context.Context, config *settings.Cache) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
