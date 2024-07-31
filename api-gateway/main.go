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

	cfg, err := config.New()
	if err != nil {
		logger.Fatalln("failed to configure project")
	}

	h := handler.New(handler.WithHTTPHandler())

	srv, err := server.New(server.WithHTTPServer(h.Mux, cfg.ApiGateway.Port))
	if err != nil {
		logger.Errorln("failed to create server")
		return
	}

	if err := srv.Start(); err != nil {
		logger.Errorln("failed to start server")
		return
	}

	logger.Infof("server is running on port %s, swagger is at /swagger/index.html\n", cfg.ApiGateway.Port)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-shutdown
	logger.Infoln("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorln("failed to stop server")
		return
	}

	logger.Infoln("server stopped")

}
