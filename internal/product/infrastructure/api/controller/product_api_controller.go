package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	productController "github.com/mathefer/tc-fiap-product/internal/product/controller"
	"github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/dto"
)

type productApiController struct {
	controller productController.ProductController
}

func NewProductController(controller productController.ProductController) *productApiController {
	return &productApiController{
		controller: controller,
	}
}

func (c *productApiController) RegisterRoutes(r chi.Router) {
	prefix := "/v1/product"
	r.Get(prefix, c.Get)
	r.Post(prefix, c.Add)
	r.Put(prefix+"/{id}", c.Update)
	r.Delete(prefix+"/{id}", c.Delete)
}

// @Summary     Get products by category
// @Description Get products by category
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       category query uint true "Category"
// @Success     200  {object} dto.GetProductResponseDto
// @Router      /v1/product [get]
// @Description Category values: 1 - Lanche, 2 - Acompanhamento, 3 - Bebida, 4 - Sobremesa
func (h *productApiController) Get(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	if category == "" {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	categoryInt, err := strconv.ParseUint(category, 10, 64)
	if err != nil {
		http.Error(w, "Invalid category parameter", http.StatusBadRequest)
		return
	}

	products, err := h.controller.Get(uint(categoryInt))

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// @Summary     Add product
// @Description Add product
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       body body dto.AddProductRequestDto true "Body"
// @Success     201
// @Router      /v1/product [post]
// @Description Category values: 1 - Lanche, 2 - Acompanhamento, 3 - Bebida, 4 - Sobremesa
func (h *productApiController) Add(w http.ResponseWriter, r *http.Request) {
	var productRequest dto.AddProductRequestDto

	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	err := h.controller.Add(&productRequest)

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary     Update product
// @Description Update product
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path uint true "Id"
// @Param       name body dto.UpdateProductRequestDto true "Name"
// @Success     200
// @Router      /v1/product/{id} [put]
// @Description Category values: 1 - Lanche, 2 - Acompanhamento, 3 - Bebida, 4 - Sobremesa
func (h *productApiController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	var productRequest dto.UpdateProductRequestDto

	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	err = h.controller.Update(id, &productRequest)

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary     Delete product
// @Description Delete product
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path uint true "Id"
// @Success     204
// @Router      /v1/product/{id} [delete]
func (h *productApiController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
	}

	err = h.controller.Delete(id)

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func getIDFromPath(r *http.Request) (uint, error) {
	vars := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(vars, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
