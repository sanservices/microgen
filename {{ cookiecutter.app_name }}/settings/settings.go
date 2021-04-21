package settings

import (
	"context"
	logger "github.com/sanservices/apilogger/v2"
	"github.com/spf13/viper"
)

const (
	defaultConfigName = "settings" // config file name without extension
)

// Settings struct
type Settings struct {
	Service service  `yaml:"service"`
	DB      database `yaml:"database"`
}

type service struct {
	Name       string `yaml:"service_name"`
	PathPrefix string `yaml:"path_prefix"`
	Version    string `yaml:"version"`
	Port       int    `yaml:"port"`
	Debug      bool   `yaml:"debug"`
}

type database struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// New Settings constructor
func New(ctx context.Context) *Settings {
	cfg, err := Get(defaultConfigName)
	if err != nil {
		logger.Fatalf(ctx, logger.LogCatReadConfig, "Fatal error config file: %s \n", err)
	}
	return cfg
}

// Get returns main configuration object
// configName param should be config file name without extension (file should be located in ./settings/)
func Get(configName string) (*Settings, error) {
	// Create viper config instance
	confer := viper.New()
	// We use yaml so there it is
	confer.SetConfigType("yaml")
	confer.SetConfigName(configName)
	// Configs are located under ./settings folder
	confer.AddConfigPath("./settings/")
	err := confer.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		return nil, err
	}

	// Fill in the struct
	return &Settings{
		Service: service{
			Name:       confer.GetString("service.name"),
			PathPrefix: confer.GetString("service.path_prefix"),
			Version:    confer.GetString("service.version"),
			Port:       confer.GetInt("service.port"),
			Debug:      confer.GetBool("service.debug"),
		},
		DB: database{
			Engine:   confer.GetString("database.engine"),
			Host:     confer.GetString("database.host"),
			Name:     confer.GetString("database.name"),
			Port:     confer.GetInt("database.port"),
			User:     confer.GetString("database.user"),
			Password: confer.GetString("database.password"),
		},
	}, nil
}
