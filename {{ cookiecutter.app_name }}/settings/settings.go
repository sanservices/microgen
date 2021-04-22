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
	Service service  `mapstructure:"service"`
	DB      database `mapstructure:"database"`
}

type service struct {
	Name       string `mapstructure:"service_name"`
	PathPrefix string `mapstructure:"path_prefix"`
	Version    string `mapstructure:"version"`
	Port       int    `mapstructure:"port"`
	Debug      bool   `mapstructure:"debug"`
}

type database struct {
	Engine   string `mapstructure:"engine"`
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
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
	settings := Settings{}

	err = confer.Unmarshal(&settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
