package app

import (
	"github/karchx/coffee-system/cmd/product/config"
	productsUC "github/karchx/coffee-system/internal/product/usecases/products"
)

type App struct {
	Cfg *config.Config
	UC  productsUC.UseCase
}

func New(cfg *config.Config, uc productsUC.UseCase) *App {
	return &App{
		Cfg: cfg,
		UC:  uc,
	}
}
