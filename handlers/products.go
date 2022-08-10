package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"example.com/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("In Put")
		path := regexp.MustCompile(`/([0-9]+)`)
		g := path.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Printf("Submatch: %#v", g)
		if len(g) != 1 {
			http.Error(rw, "Invalid Uri", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid Uri 2", http.StatusBadRequest)
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		p.l.Println("Got Id", id)
		p.updateProducts(id, rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle post products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall data", http.StatusBadRequest)
	}

	data.AddProducts(product)
	p.l.Printf("prod: %#v", product)
}

/*
* This method will decode the request json
to the product
*/
func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle update products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall data", http.StatusBadRequest)
	}

	err = data.UpdateProducts(id, product)

	if err == data.ErrProductNotFound {
		p.l.Println("Error - 1 %#v", err)
		http.Error(rw, "Product Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		p.l.Println("Error - 2 %#v", err)
		http.Error(rw, "Another problem", http.StatusNotFound)
		return
	}
	p.l.Printf("update prod: %#v", product)
}
