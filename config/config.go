package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
}

func Default() Config {
	return Config{}
}

func Init() (*Config, error) {
	config := Default()

	err := envconfig.Process("project", &config)
	if err != nil {
		return nil, errors.New("failed to read environnement variable")
	}

	return &config, nil
}
