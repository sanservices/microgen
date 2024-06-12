package config

import (
	"context"
	"os"
	{% if cookiecutter.use_cache != 'n' %}"time"{% endif %}

	{% if cookiecutter.use_database != 'n' %}"github.com/sanservices/kit/database"{% endif %}
	"gopkg.in/yaml.v2"
)

const (
	fPath = "settings.yml"
)

type Settings struct {
	Service  Service                   `yaml:"Service"`
	{% if cookiecutter.use_database == 'y' %}Database  database.DatabaseConfig                  `yaml:"Database"`{% endif %}
	{% if cookiecutter.use_cache == 'y' %}Cache    CacheConfig                     `yaml:"Cache"`{% endif %}
	{% if cookiecutter.use_sftp == 'y' %}Sftp    SftpConfig                `yaml:"Sftp"`{% endif %}
}

type Service struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
	// TODO: Add debug flag
}

{% if cookiecutter.use_cache == 'y' %}
type CacheConfig struct {
	Enabled           bool                 `yaml:"enabled"`
	ExpirationMinutes time.Duration        `yaml:"expiration_minutes"`
	PurgeMinutes      time.Duration        `yaml:"purge_minutes"`
	RedisConfig       database.RedisConfig `yaml:"Redis"`
}
{% endif %}

{% if cookiecutter.use_sftp == 'y' %}
type SftpConfig struct { 
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}
{% endif %}

func New(ctx context.Context) (*Settings, error) {
	settings := &Settings{}

	//Read settings file
	cf, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(cf, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
