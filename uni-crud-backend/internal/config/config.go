package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	JWTSecret string
}

func InitializeConfig() (Config, error) {
	file, err := os.Open("../config.yaml")

	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
