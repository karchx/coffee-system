// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github/karchx/coffee-system/cmd/product/config"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, grpcServer *grpc.Server) (*App, error) {
	app := New(cfg)
	return app, nil
}
