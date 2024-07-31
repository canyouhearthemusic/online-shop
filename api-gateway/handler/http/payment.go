package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PaymentHandler struct {
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Post("/", h.create)

	r.Get("/search", h.search)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

	return r
}

func (h *PaymentHandler) list(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentHandler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentHandler) get(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentHandler) search(w http.ResponseWriter, r *http.Request) {

}
