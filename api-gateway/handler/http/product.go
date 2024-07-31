package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) Routes() chi.Router {
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

func (h *ProductHandler) list(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) get(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) search(w http.ResponseWriter, r *http.Request) {

}
