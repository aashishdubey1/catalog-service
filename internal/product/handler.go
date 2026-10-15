package product

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAllProudts)
	r.Post("/", h.CreateProduct)
	r.Get("/{id}", h.GetProductById)

	return r
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not imlemented yet"))
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("not imlemented yet"))
}

func (h *ProductHandler) GetAllProudts(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("not imlemented yet"))
}
