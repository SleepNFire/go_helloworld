package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

type MySql struct {
	Adress   string
	User     string
	Password string
	Database string
}

type Config struct {
	MySql MySql
}

func Default() Config {
	return Config{
		MySql: MySql{
			Adress:   "127.0.0.1:3306",
			User:     "user",
			Password: "password",
			Database: "PROJECT",
		},
	}
}

func Init() (*Config, error) {
	config := Default()

	err := envconfig.Process("project", &config)
	if err != nil {
		return nil, errors.New("failed to read environnement variable")
	}

	return &config, nil
}
