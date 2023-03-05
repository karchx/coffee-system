package main

import (
	"context"
	"fmt"
	"github/karchx/coffee-system/cmd/product/config"
	"github/karchx/coffee-system/internal/product/app"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	slog.Info("⚡ init app", "name", cfg.Name, "version", cfg.Version)

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	_, err = app.InitApp(cfg, server)
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
	}

	// gRPC server.
	address := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.PORT)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
	}

	slog.Info(" start server...", "address", address)

	defer func() {
		if err1 := l.Close(); err1 != nil {
			slog.Error("failed to close", err1, "network", network, "address", address)
		}
	}()

	err = server.Serve(l)
	if err != nil {
		slog.Error("failed start gRPC server", err, "network", network, "address", address)
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done", done)
	}
}
