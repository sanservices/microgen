package settings

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
	"{{ cookiecutter.module_name }}/goutils/tlsutil"
)

type HeaderKey string

const (
	defaultFilePath = "settings.yml"

	// EmailKey key to lookup email value in request context
	EmailKey HeaderKey = "email"

	// UsernNameKey key to lookup username value in request context
	UsernNameKey HeaderKey = "preferred_username"

	// defines date formats for layouts
	SimpleDateLayout = "2006-01-02"

	// EnrolledProgramsMaxQuantity maximun amount of programs
	// an user can be enrolled at once.
	EnrolledProgramsMaxQuantity = 5
)

var ErrNoFile error = errors.New("file not found")
var ErrParsingFile error = errors.New("unable to parse file")

type Settings struct {
	Service     Service         `yaml:"Service"`
	Cache       Cache           `yaml:"Cache"`
	DB          Database        `yaml:"Database"`
	LegacyDB    Database        `yaml:"LegacyDatabase"`
	Kafka       Kafka           `yaml:"Kafka"`
	AuthService ExternalService `yaml:"AuthService"`
}

type Service struct {
	Name       string `yaml:"name"`
	PathPrefix string `yaml:"path_prefix"`
	Version    string `yaml:"version"`
	Port       int    `yaml:"port"`
	Debug      bool   `yaml:"debug"`
}

type Cache struct {
	Enabled           bool          `yaml:"enabled"`
	Addr              string        `yaml:"addr"`
	ExpirationMinutes time.Duration `yaml:"expiration_minutes"`
	PurgeMinutes      time.Duration `yaml:"purge_minutes"`
}

type Database struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Kafka struct {
	Enabled         bool        `yaml:"enabled"`
	Version         string      `yaml:"version"`
	ConsumeOnly     bool        `yaml:"consume_only"`
	ConsumerGroupID string      `yaml:"consumer_group_id"`
	FromOldest      bool        `yaml:"from_oldest"`
	IsTLS           bool        `yaml:"is_tls"`
	SchemaReg       string      `yaml:"schema_registration"`
	Brokers         []string    `yaml:"brokers"`
	TLS             tlsutil.TLS `yaml:"TLS"`
}

type ExternalService struct {
	BaseURL string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
}

// GetEnv returns env value, if empty, returns fallback value
func GetEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

//Get returns main configuration object
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

func New(ctx context.Context) (*Settings, error) {
	return Get(ctx, defaultFilePath)
}
