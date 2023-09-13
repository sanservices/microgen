package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sanservices/kit/database"
	"{{ cookiecutter.module_name }}/config"
)

type Cache struct {
	client  *redis.Client
	expTime time.Duration
	enabled bool
}

var ErrCacheDisabled error = errors.New("cache is not enabled")

// New returns a Cache instance
func New(ctx context.Context, config *config.Config) (*Cache, error) {

	var rdb *redis.Client
	var err error

	if config.Cache.Enabled {
		rdb, err = database.CreateRedisConnection(ctx, config.Cache.RedisConfig)
		if err != nil {
			return nil, err
		}
	}

	c := &Cache{
		client:  rdb,
		expTime: config.Cache.ExpirationMinutes * time.Minute,
		enabled: config.Cache.Enabled,
	}

	return c, nil
}

// Set stores the value in redis with given key
func (c *Cache) Set(ctx context.Context, key string, v interface{}) error {
	if !c.enabled {
		return ErrCacheDisabled
	}

	bytes, _ := json.Marshal(v)

	return c.client.Set(ctx, key, bytes, c.expTime).Err()
}

// Get retreives the value of the given key, if it doesn't find it will return redis: nil
func (c *Cache) Get(ctx context.Context, key string, v interface{}) error {
	if !c.enabled {
		return ErrCacheDisabled
	}

	str, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), &v)
}

// Delete removes the value from redis of the given key
func (c *Cache) Delete(ctx context.Context, key string) error {
	if !c.enabled {
		return ErrCacheDisabled
	}

	return c.client.Del(ctx, key).Err()
}
