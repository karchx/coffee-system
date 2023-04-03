//go:build wireinject
// +build wireinject

package app

import (
	"github/karchx/coffee-system/cmd/product/config"
	"github/karchx/coffee-system/internal/product/infras/repo"
	productsUC "github/karchx/coffee-system/internal/product/usecases/products"

	"github.com/google/wire"

	"google.golang.org/grpc"
)

func InitApp(cfg *config.Config, grpcServer *grpc.Server) (*App, error) {
	panic(wire.Build(
		New,
		repo.RepositorySet,
		productsUC.UseCaseSet,
	))
}
