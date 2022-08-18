package data

import (
	"fmt"
	"time"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	// the id for the product is autogenerate
	//
	// min: 1
	ID int `json:"id"`

	// name of the product
	//
	// required: true
	// example: tea
	Name string `json:"name" validate:"required"`

	// description of the product
	Description string `json:"description"`

	// price the product
	//
	// required: true
	// example: 4.50
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// name of the product
	//
	// required: true
	// example: a-b-c
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

func UpdateProducts(id int, p *Product) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = p

	return nil
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

func findIndexByProductID(id int) int {
	fmt.Println("findProduct: Product ID: ", id)
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky coffee",
		Price:       2.45,
		SKU:         "latte12",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and Strong coffee",
		Price:       1.99,
		SKU:         "espresso12",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
