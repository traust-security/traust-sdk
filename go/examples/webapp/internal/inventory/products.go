package inventory

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// products answers the products page: core's whole catalogue.
func (h *Handler) products(w http.ResponseWriter, r *http.Request) {
	var products []traust.Product
	for product, err := range h.coreInventory.ListProductsAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		products = append(products, product)
	}
	web.Render(w, r, http.StatusOK, productList(products))
}

// product answers a product's page: the product itself and the refs core knows under it.
func (h *Handler) product(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.coreInventory.GetProduct(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	refs, err := h.coreInventory.ListProductRefs(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, productDetail(*found, refs))
}

// productsOf answers a repository page's lazy products section: the products core files the
// repository under.
func (h *Handler) productsOf(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	products, err := h.coreInventory.ListProductsOf(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, repositoryProducts(products))
}
