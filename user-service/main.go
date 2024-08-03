package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canyouhearthemusic/online-shop/common/config"
	"github.com/canyouhearthemusic/online-shop/common/server"
	"github.com/canyouhearthemusic/online-shop/user-service/internal/handler"
	"github.com/canyouhearthemusic/online-shop/user-service/internal/repository"
	onlineshop "github.com/canyouhearthemusic/online-shop/user-service/internal/service/online-shop"
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

	repo, err := repository.New(repository.WithPostgresStore(cfg.UserDB.DSN))
	if err != nil {
		logger.Fatalf("Repository couldn't be created: %s\n", err)
	}
	defer repo.Close()

	service, err := onlineshop.New(onlineshop.WithUserRepository(repo.User))
	if err != nil {
		logger.Fatalf("User online-shop service couldn't be created: %s:\n", err)
	}

	handlers := handler.New(
		handler.Dependencies{OnlineShopService: service},
		handler.WithHTTPHandler(),
	)

	srv, err := server.New(server.WithHTTPServer(handlers.Mux, cfg.User.Port))
	if err != nil {
		logger.Fatalf("Failed to create server: %s\n", err)
		return
	}

	if err := srv.Start(); err != nil {
		logger.Fatalf("Failed to start server: %s\n", err)
	}

	logger.Infof("Server is running on port %s, swagger is at /swagger/index.html\n", cfg.User.Port)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	<-shutdown
	logger.Infoln("Shutting down server")

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("Failed to stop server: %s", err)
	}

	logger.Infoln("Server stopped gracefully")
}
