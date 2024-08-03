package handler

import (
	"net/http"

	"github.com/canyouhearthemusic/online-shop/api-gateway/handler/httphandler"
	"github.com/canyouhearthemusic/online-shop/common/server/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Configuration func(h *Handler) error

type Handler struct {
	Mux *chi.Mux
}

func New(cfgs ...Configuration) Handler {
	h := Handler{}

	for _, cfg := range cfgs {
		cfg(&h)
	}

	return h
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) error {
		h.Mux = router.New()

		userHandler := httphandler.NewUserHandler()
		productHandler := httphandler.NewProductHandler()
		orderHandler := httphandler.NewOrderHandler()
		paymentHandler := httphandler.NewPaymentHandler()

		h.Mux.Route("/api/v1", func(r chi.Router) {
			r.Get("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
				render.Status(r, http.StatusOK)
				render.PlainText(w, r, "OK")
			})

			r.Mount("/users", userHandler.Routes())
			r.Mount("/products", productHandler.Routes())
			r.Mount("/orders", orderHandler.Routes())
			r.Mount("/payments", paymentHandler.Routes())
		})

		return nil
	}
}
