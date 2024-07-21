package jwt

import (
	"context"
)

type configKey struct{}

func WithConfig(ctx context.Context, cfg Config) context.Context {
	return context.WithValue(ctx, configKey{}, cfg)
}

func GetConfig(ctx context.Context) *Config {
	raw := ctx.Value(configKey{})
	cfg, ok := raw.(Config)
	if !ok {
		return nil
	}

	return &cfg
}
