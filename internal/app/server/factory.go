package server

import (
	"fmt"

	"github.com/chaewonkong/go-msa-arch/internal/app"
)

// Factory server factory
type Factory struct {
}

// NewConfig config constructor
func (f *Factory) NewConfig() any {
	return &Config{}
}

// NewApp app constructor
func (f *Factory) NewApp(config any) (app.App, error) {
	cfg, ok := config.(*Config)
	if !ok {
		return nil, fmt.Errorf("invalid config type")
	}

	return New(cfg), nil
}
