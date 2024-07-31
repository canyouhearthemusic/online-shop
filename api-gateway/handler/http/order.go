package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (h *OrderHandler) Routes() chi.Router {
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

func (h *OrderHandler) list(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) get(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) search(w http.ResponseWriter, r *http.Request) {

}
