package utils

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		ExitError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		ExitError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		ExitError(err)
	}
}

func ReadConfig() Config {
	var config Config

	readFile(&config)
	readEnv(&config)

	return config
}