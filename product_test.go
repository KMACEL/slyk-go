package slyk

import "testing"

func TestGetProducts(t *testing.T) {
	tst := "TestGetProducts"

	returnValue, err := getClient().GetProducts()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetProductWithID(t *testing.T) {
	tst := "TestGetProductWithID"

	returnValue, err := getClient().GetProductWithID("")

	ReturnAndError(t, tst, returnValue, err)
}
