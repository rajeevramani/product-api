package data

import "testing"

func TestValdidateProduct(t *testing.T) {
	prod := &Product{
		Name: "Tea",
		Price: 3.40,
		SKU: "a-b-c",
	}

	err := prod.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
