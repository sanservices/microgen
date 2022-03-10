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
	"github.com/sanservices/kit/database"
)

var (
	// ErrInvalidDBEngine database engine is not supported
	ErrInvalidDBEngine = errors.New("unsoported or missing database engine")
)

func New(config *settings.Settings) (*sqlx.DB, error) {

	switch config.DB.Engine {
	case "mysql":
		return CreateMySqlConnection(&config.DB)

	case "sqlite":
		return CreateSqliteConnection(&config.DB)

	default:
		return nil, ErrInvalidDBEngine
	}
}

func CreateMySqlConnection(dbConfig *database.DatabaseConfig) (*sqlx.DB, error) {

	var connectionString string
	var db *sqlx.DB
	var err error

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

func CreateSqliteConnection(config *database.DatabaseConfig) (*sqlx.DB, error) {
	log.Println("Connecting to database...")
	source := fmt.Sprintf("./%s.db", config.Name)

	db, err := sqlx.Connect("sqlite3", source)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}

func CreateRedisConnection(ctx context.Context, config *database.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.DB,  // use default DB
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}