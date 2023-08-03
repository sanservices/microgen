package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"{{ cookiecutter.module_name }}/config"
)

func New(ctx context.Context, config *config.Config) (*redis.Client, error) {

	options := &redis.Options{
		Addr:     config.Cache.Address,
		Password: config.Cache.Password,
		DB:       config.Cache.DB,
	}

	client := redis.NewClient(options)

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
