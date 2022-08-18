package handlers

import (
	"net/http"

	"example.com/product-api/data"
)

// swagger:route GET /products products listProducts
//
// returns a lists of products.
// Responses:
//   200: productResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := data.ToJSON(lp, rw)
	if err != nil {
		http.Error(rw, "Invalid data", http.StatusInternalServerError)
	}
}
