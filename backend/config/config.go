package config

import (
	"log/slog"
	"os"
	"sync"
	"true-kw/pkg/db/client/postgresql"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "config/config.yml"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	Storage StorageConfig `yaml:"storage"`
}

type AppConfig struct {
	LogLevel string `yaml:"log_level" env-default:"debug"`
	Prod     bool   `yaml:"prod" env-default:"false"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"8080"`
}

type StorageConfig struct {
	PostgresConfig postgresql.StorageConfig `yaml:"postgres"`
}

var instance *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		instance = &Config{}
		readErr := cleanenv.ReadConfig(configPath, instance)
		if readErr != nil {
			description, descrErr := cleanenv.GetDescription(instance, nil)
			if descrErr != nil {
				panic(descrErr)
			}
			slog.Info(description)
			slog.Error(
				"failed to read config",
				slog.String("err", readErr.Error()),
				slog.String("path", configPath),
			)
			os.Exit(1)
		}
	})
	return instance
}
