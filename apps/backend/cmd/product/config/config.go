package config

import (
	"fmt"
	"log"
	"os"

	configs "github/karchx/coffee-system/pkg/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	configs.App  `yaml:"app"`
	configs.HTTP `yaml:"http"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
