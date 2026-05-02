package config

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"

	"fintech-proj/util/validator"
)

var once sync.Once

func loaddotenv(filename string) {
	once.Do(func() {
		// Ignore errors - these files are optional and not even present when the app is started from inside a Docker
		// container.
		_ = godotenv.Load(filename)
		_ = godotenv.Load(".env.common")
	})
}

type Config struct {
	Port     int        `env:"PORT" validate:"required,min=1,max=65535"`
	LogLevel slog.Level `env:"LOG_LEVEL" validate:"required"`
}

func NewConfig(path string) (Config, error) {
	loaddotenv(path)

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("applying env from file: %w", err)
	}

	if err := validator.Validator.Struct(&cfg); err != nil {
		return Config{}, fmt.Errorf("validating config: %w", err)
	}

	return cfg, nil
}
