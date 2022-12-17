package config

import (
	"log"
	"os"

	"github.com/worldwidepaniel/uni-crud-project/internal/structs"
	"gopkg.in/yaml.v3"
)

func InitializeConfig() structs.Config {
	file, err := os.Open("./config.yaml")

	if err != nil {
		log.Fatal(err)
		return structs.Config{}
	}

	defer file.Close()

	var cfg structs.Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
		return structs.Config{}
	}
	return cfg
}
