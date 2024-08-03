package httphandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Routes() chi.Router {
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

func (h *UserHandler) list(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) search(w http.ResponseWriter, r *http.Request) {

}
