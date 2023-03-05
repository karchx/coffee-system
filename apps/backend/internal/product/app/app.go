package app

import (
	"github/karchx/coffee-system/cmd/product/config"
)

type App struct {
	Cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		Cfg: cfg,
	}
}
