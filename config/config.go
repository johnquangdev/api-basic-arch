package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB struct {
		Host     string `envconfig:"DB_HOST"`
		User     string `envconfig:"DB_USER"`
		Password string `envconfig:"DB_PASSWORD"`
		Name     string `envconfig:"DB_NAME"`
		Port     int    `envconfig:"DB_PORT"`
	}
	JWT struct {
		SecretKey     string `envconfig:"SECRET_KEY"`
		TokenTimeLife int    `envconfig:"TOKEN_TIME_LIFE"`
	}
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("load config error: %v", err)
	}
	return cfg, nil
}
