package httphandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/canyouhearthemusic/online-shop/common/server/response"
	"github.com/canyouhearthemusic/online-shop/user-service/internal/domain/user"
	onlineshop "github.com/canyouhearthemusic/online-shop/user-service/internal/service/online-shop"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *onlineshop.Service
}

func NewUserHandler(service *onlineshop.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Routes() *chi.Mux {
	r := chi.NewMux()

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
	data, err := h.service.ListUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.OK(w, r, data)
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	req := user.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NotFound(w, r, err)
		return
	}

	if err := req.Validate(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)

		return
	}

	err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.Created(w, r, "user has been created successfully!", "OK")
}

func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.OK(w, r, data)
}

func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := user.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NotFound(w, r, err)
		return
	}

	if errs := req.Validate(); errs != nil {
		response.BadRequests(w, r, errs)
		return
	}

	err := h.service.UpdateUser(r.Context(), id, &req)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			response.NotFound(w, r, err)
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			response.NotFound(w, r, err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) search(w http.ResponseWriter, r *http.Request) {
	req := &user.Request{
		Name:  r.URL.Query().Get("name"),
		Email: r.URL.Query().Get("email"),
	}

	res, err := h.service.SearchUsers(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
