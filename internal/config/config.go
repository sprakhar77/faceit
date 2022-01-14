package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Application contains all the required application config
type Application struct {
	ID          string    `mapstructure:"APPLICATION_ID"`
	Name        string    `mapstructure:"APPLICATION_NAME"`
	Version     string    `mapstructure:"VERSION"`
	Environment string    `mapstructure:"ENVIRONMENT"`
	Log         Log       `mapstructure:"LOG"`
	Server      Server    `mapstructure:"Server"`
	Database    Database  `mapstructure:"DATABASE"`
	Publisher   Publisher `mapstructure:"PUBLISHER"`
}

type Server struct {
	Host            string        `mapstructure:"HOST"`
	Port            string        `mapstructure:"PORT"`
	ShutdownTimeout time.Duration `mapstructure:"SHUTDOWN_TIMEOUT"`
}

type Database struct {
	Name         string        `mapstructure:"NAME"`
	User         string        `mapstructure:"USER"`
	Password     string        `mapstructure:"PASSWORD" json:"-"`
	Host         string        `mapstructure:"HOST"`
	Port         string        `mapstructure:"PORT"`
	SSLMode      string        `mapstructure:"SSL_MODE"`
	QueryTimeout time.Duration `mapstructure:"QUERY_TIMEOUT"`
}

type Log struct {
	Level     string `mapstructure:"LEVEL"`
	Formatter string `mapstructure:"FORMATTER"`
}

type KafkaWriterConfig struct {
	URL       string `mapstructure:"URL"`
	EventType string `mapstructure:"EVENT_TYPE"`
	Enabled   bool   `mapstructure:"ENABLED"`
}

type Publisher struct {
	User KafkaWriterConfig `mapstructure:"MATERIALIZED_PRICE"`
}

// Load method reads the configuration from a file and returns it as a viper.Application pointer
func Load() (*Application, error) {
	// read the config.yaml/mapstructure file into viper
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// move config data into struct
	app := Application{}
	return &app, viper.Unmarshal(&app)
}
