package config

import (
	"context"
	"os"

	{% if cookiecutter.use_cache == 'y' %}"time"{% endif %}

	{% if cookiecutter.use_database == 'y' or cookiecutter.use_cache == 'y' %}"github.com/sanservices/kit/database"{% endif %}
	{% if cookiecutter.use_kafka == 'y' %}"github.com/sanservices/kit/kafkalistener"{% endif %}
	"gopkg.in/yaml.v2"
)

const (
	// defaultSettingsPath is used when the SETTINGS_PATH env var is not set.
	defaultSettingsPath = "settings.yml"
)

type Settings struct {
	Service  Service                   `yaml:"Service"`
	GRPC     GRPC                      `yaml:"GRPC"`
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

type GRPC struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    string `yaml:"port"`
}

{% if cookiecutter.use_cache == 'y' %}
type Cache struct {
	Enabled           bool                 `yaml:"enabled"`
	ExpirationMinutes time.Duration        `yaml:"expiration_minutes"`
	PurgeMinutes      time.Duration        `yaml:"purge_minutes"`
	RedisConfig       database.RedisConfig `yaml:"Redis"`
}
{% endif %}

func New(ctx context.Context) (*Settings, error) {
	settings := &Settings{}

	// Read settings file (path overridable via the SETTINGS_PATH env var)
	path := os.Getenv("SETTINGS_PATH")
	if path == "" {
		path = defaultSettingsPath
	}

	cf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(cf, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
