package settings

import (
	"context"
	"errors"
	"io/ioutil"
	"time"

	"github.com/sanservices/kit/database"
	"gopkg.in/yaml.v2"
)

type HeaderKey string

const (
	defaultFilePath = "settings.yml"
)

var (
	ErrNoFile      error = errors.New("file not found")
	ErrParsingFile error = errors.New("unable to parse file")
)

type Settings struct {
	Service Service                 `yaml:"Service"`
	Cache   Cache                   `yaml:"Cache"`
	DB      database.DatabaseConfig `yaml:"Database"`
}

type Service struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
	Debug   bool   `yaml:"debug"`
}

type Cache struct {
	Enabled           bool                 `yaml:"enabled"`
	ExpirationMinutes time.Duration        `yaml:"expiration_minutes"`
	PurgeMinutes      time.Duration        `yaml:"purge_minutes"`
	Redis             database.RedisConfig `yaml:"redis"`
}

func GetDefault(ctx context.Context) (*Settings, error) {
	return Get(ctx, defaultFilePath)
}

func Get(ctx context.Context, filePath string) (*Settings, error) {
	var err error
	var confFile []byte

	config := &Settings{}

	confFile, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, ErrNoFile
	}

	//if file exists use its variables
	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		return nil, ErrParsingFile
	}

	return config, nil
}
