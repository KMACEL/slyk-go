package slyk

import "testing"

func TestGetTaxRates(t *testing.T) {
	tst := "TestGetTaxRates"

	returnValue, err := getClient().GetTaxRates()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetTaxRateWithID(t *testing.T) {
	tst := "TestGetTaxRateWithID"

	returnValue, err := getClient().GetTaxRateWithID("19e0ff97-3588-410c-a225-efa5f9f26493")

	ReturnAndError(t, tst, returnValue, err)
}
