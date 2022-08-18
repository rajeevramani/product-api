package handlers

import (
	"net/http"

	"example.com/product-api/data"
)

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	product := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("[DEBUG] updating product with ID: ", product.ID)

	errp := data.UpdateProducts(id, product)

	if errp == data.ErrProductNotFound {
		p.l.Println("[ERROR] Product nof found", errp)
		http.Error(rw, "Product Not found", http.StatusNotFound)
		// data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return

	}
	p.l.Printf("update prod: %#v", product)
}
