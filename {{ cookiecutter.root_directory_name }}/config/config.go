package config

import (
	"context"
	"os"

	{% if cookiecutter.use_cache == 'y' %}"time"{% endif %}

	{% if cookiecutter.use_database != 'n' %}"github.com/sanservices/kit/database"{% elif cookiecutter.use_cache != 'n' %}"github.com/sanservices/kit/database"{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}"github.com/sanservices/kit/kafkalistener"{% endif %}
	"gopkg.in/yaml.v2"
)

const (
	fPath = "config.yml"
)

type Config struct {
	Service  Service                   `yaml:"Service"`
	{% if cookiecutter.use_database == 'y' %}Database  database.DatabaseConfig                  `yaml:"Database"`{% endif %}
	{% if cookiecutter.use_cache == 'y' %}Cache    Cache                     `yaml:"Cache"`{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}Kafka    kafkalistener.KafkaConfig `yaml:"Kafka"`{% endif %}
}

type Service struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
	// TODO: Add debug flag
}

{% if cookiecutter.use_cache == 'y' %}
type Cache struct {
	Enabled           bool                 `yaml:"enabled"`
	ExpirationMinutes time.Duration        `yaml:"expiration_minutes"`
	PurgeMinutes      time.Duration        `yaml:"purge_minutes"`
	RedisConfig       database.RedisConfig `yaml:"Redis"`
}
{% endif %}

func New(ctx context.Context) (*Config, error) {
	config := &Config{}

	//Read config file
	cf, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(cf, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
