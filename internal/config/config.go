package config

import (
	"context"
	"log"
	"os"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	API_KEY  string `env:"API_KEY, required"`
	Filepath string
}

func Load(ctx context.Context) (*Config, error) {
	var cfg Config

	err := envconfig.Process(ctx, &cfg)
	if err != nil {
		return &Config{}, err
	}

	file, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	file += "\\crypto-storage\\storage.json"

	cfg.Filepath = file

	return &cfg, nil
}
