package handler

import (
	"github.com/canyouhearthemusic/online-shop/common/server/router"
	"github.com/canyouhearthemusic/online-shop/user-service/internal/handler/httphandler"
	onlineshop "github.com/canyouhearthemusic/online-shop/user-service/internal/service/online-shop"
	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	OnlineShopService *onlineshop.Service
}

type Handler struct {
	deps Dependencies

	Mux *chi.Mux
}

type Configuration func(h *Handler) error

func New(deps Dependencies, configs ...Configuration) Handler {
	h := Handler{
		deps: deps,
	}

	for _, cfg := range configs {
		cfg(&h)
	}

	return h
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) error {
		h.Mux = router.New()

		userHandler := httphandler.NewUserHandler(h.deps.OnlineShopService)

		h.Mux.Route("/api/v1", func(r chi.Router) {
			r.Mount("/users", userHandler.Routes())
		})

		return nil
	}
}
