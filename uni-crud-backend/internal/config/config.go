package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	JWTSecret string `yaml:"JWTSecret"`
	DBPath    string `yaml:"DBPath"`
	Port      string `yaml:"Port"`
}

func (c Config) InitializeConfig() Config {
	file, err := os.Open("./config.yaml")

	if err != nil {
		log.Fatal(err)
		return Config{}
	}

	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
		return Config{}
	}
	return cfg
}
