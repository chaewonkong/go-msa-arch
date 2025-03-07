package server

import (
	"fmt"

	"github.com/chaewonkong/go-msa-arch/internal/app"
)

type Factory struct {
}

func (f *Factory) NewConfig() any {
	return &Config{}
}

func (f *Factory) NewApp(config any) (app.App, error) {
	cfg, ok := config.(*Config)
	if !ok {
		return nil, fmt.Errorf("invalid config type")
	}

	return New(cfg), nil
}
