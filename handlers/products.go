// Package classification Petstore API.
//
// the purpose of this application is to provide an application
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /pets pets users listPets
//
// Lists pets filtered by some parameters.
// Responses:
//   default: genericError
//   200: productResponse
//   500: InternalServerError

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Invalid data", http.StatusInternalServerError)
	}
}

/*
* This method will decode the request json
to the product
*/
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle post products")

	// product := &data.Product{}
	// err := product.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshall data", http.StatusBadRequest)
	// }

	// prod := r.Context().Value(KeyProduct{})(data.Product)
	// prod := r.Context().Value("KeyProduct")(data.Product)

	pr := r.Context().Value(ContextUserKey).(data.Product)
	data.AddProducts(&pr)

	// product := r.Context().Value(KeyProduct{})

	// data.AddProducts(&product)
	p.l.Printf("prod: %#v", pr)
}

/*
* This method will decode the request json
to the product
*/
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle update products")
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "invalid id provided", http.StatusBadRequest)
	}

	// product := &data.Product{}
	// errp := product.FromJSON(r.Body)
	// if errp != nil {
	// 	http.Error(rw, "Unable to unmarshall data", http.StatusBadRequest)
	// 	return
	// }
	// prod := r.Context().Value("KeyProduct").(data.Product)
	product := r.Context().Value(ContextUserKey).(data.Product)

	errp := data.UpdateProducts(id, &product)

	if errp == data.ErrProductNotFound {
		p.l.Printf("Error - 1 %#v", errp)
		http.Error(rw, "Product Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		p.l.Printf("Error - 2 %#v", errp)
		http.Error(rw, "Another problem", http.StatusNotFound)
		return
	}
	p.l.Printf("update prod: %#v", product)
}

/*
adding middleware to handle common functionality between POST and PUT
where we get something in the body of the request
*/

// type KeyProduct struct {
// }

type ContextKey string

const ContextUserKey ContextKey = "product"

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}
		errp := product.FromJSON(r.Body)
		if errp != nil {
			http.Error(rw, "Unable to unmarshall data", http.StatusBadRequest)
			return
		}

		// adding validation

		err := product.Validate()
		if err != nil {
			p.l.Println("[Error] validation product: ", err)
			http.Error(rw, fmt.Sprintf("Validation failed: %s", err), http.StatusBadRequest)
			return
		}
		// ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		ctx := context.WithValue(r.Context(), ContextUserKey, product)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
