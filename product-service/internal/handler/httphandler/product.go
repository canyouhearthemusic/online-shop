package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/canyouhearthemusic/online-shop/common/server/response"
	"github.com/canyouhearthemusic/online-shop/product-service/internal/domain/product"
	onlineshop "github.com/canyouhearthemusic/online-shop/product-service/internal/service/online-shop"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	service *onlineshop.Service
}

func NewProductHandler(service *onlineshop.Service) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Routes() *chi.Mux {
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
	data, err := h.service.ListProducts(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.OK(w, r, data)
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	req := product.Entity{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, r, err, "Bad Request")
		return
	}

	if err := req.Validate(); err != nil {
		response.BadRequest(w, r, err, "Bad Request")

		return
	}

	err := h.service.CreateProduct(r.Context(), &req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.Created(w, r, "user has been created successfully!", "OK")
}

func (h *ProductHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, err := h.service.GetProduct(r.Context(), id)
	if err != nil {
		logrus.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.OK(w, r, data)
}

func (h *ProductHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := product.Entity{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NotFound(w, r, err)
		return
	}

	if errs := req.Validate(); errs != nil {
		response.BadRequests(w, r, errs)
		return
	}

	err := h.service.UpdateProduct(r.Context(), id, &req)
	if err != nil {
		response.NotFound(w, r, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteProduct(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *ProductHandler) search(w http.ResponseWriter, r *http.Request) {
	req := &product.Entity{
		Title:    r.URL.Query().Get("title"),
		Category: r.URL.Query().Get("category"),
	}

	res, err := h.service.SearchProducts(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
