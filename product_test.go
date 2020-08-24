package slyk

import "testing"

func TestGetProducts(t *testing.T) {
	tst := "TestGetProducts"

	returnValue, err := getClient().GetProducts()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetProductWithID(t *testing.T) {
	tst := "TestGetProductWithID"

	returnValue, err := getClient().GetProductWithID("bfbd599a-294b-4e52-989f-ec8d6b281950")

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateProducts(t *testing.T) {
	tst := "TestCreateProducts"

	returnValue, err := getClient().CreateProducts(&CreateProductsDataBody{
		CategoryID: "27d3d24a-257b-4a88-a8eb-bb8b6121f317",
		Price:      "10.2",
		Name:       "Mert",
	})

	ReturnAndError(t, tst, returnValue, err)
}
