package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	API_KEY string `env:"API_KEY, required"`
}

func Load(ctx context.Context) (*Config, error) {
	var cfg Config

	err := envconfig.Process(ctx, &cfg)
	if err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}
