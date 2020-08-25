package slyk

import "testing"

func TestGetTaxRates(t *testing.T) {
	tst := "TestGetTaxRates"

	returnValue, err := getClient().GetTaxRates()

	ReturnAndError(t, tst, returnValue, err)
}
