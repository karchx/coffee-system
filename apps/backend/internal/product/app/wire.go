//go:build wireinject
// +build wireinject

package app

import (
	"github/karchx/coffee-system/cmd/product/config"

	"github.com/google/wire"

	"google.golang.org/grpc"
)


func InitApp(cfg *config.Config, grpcServer *grpc.Server) (*App, error) {
  panic(wire.Build(New))
}
