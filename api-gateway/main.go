package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canyouhearthemusic/online-shop/api-gateway/handler"
	"github.com/canyouhearthemusic/online-shop/common/config"
	"github.com/canyouhearthemusic/online-shop/common/server"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	logger := logrus.New().WithContext(ctx)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Config couldn't be loaded: %s\n", err)
		return
	}

	h := handler.New(handler.WithHTTPHandler())

	srv, err := server.New(server.WithHTTPServer(h.Mux, cfg.ApiGateway.Port))
	if err != nil {
		logger.Fatalf("Failed to create server: %s\n", err)
		return
	}

	if err := srv.Start(); err != nil {
		logger.Fatalf("Failed to start server: %s\n", err)
	}

	logger.Infof("Server is running on port %s, swagger is at /swagger/index.html\n", cfg.ApiGateway.Port)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-shutdown
	logger.Infoln("Shutting down server")

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("Failed to stop server: %s", err)
	}

	logger.Infoln("Server stopped gracefully")
}
