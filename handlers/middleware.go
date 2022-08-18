package handlers

import (
	"context"
	"net/http"

	"example.com/product-api/data"
)

type ContextKey string

const ContextUserKey ContextKey = "product"

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		p.l.Printf("[Infor] product %#v ", prod)
		p.l.Println("")

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

// func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

// 		product := &data.Product{}
// 		errp := data.FromJSON(product, r.Body)
// 		if errp != nil {

// 			p.l.Println("[Error] getting payloag", errp)

// 			http.Error(rw, "Unable to unmarshall data", http.StatusBadRequest)
// 			return
// 		}

// 		// adding validation

// 		err := p.v.Validate(product)
// 		if len(err) != 0 {
// 			p.l.Println("[Error] validation product: ", err)
// 			http.Error(rw, fmt.Sprintf("Validation failed: %s", err), http.StatusBadRequest)
// 			return
// 		}
// 		// ctx := context.WithValue(r.Context(), KeyProduct{}, product)
// 		ctx := context.WithValue(r.Context(), ContextUserKey, product)
// 		r = r.WithContext(ctx)

// 		next.ServeHTTP(rw, r)
// 	})
// }
