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

func TestUpdateTaxRate(t *testing.T) {
	tst := "TestUpdateTaxRate"

	returnValue, err := getClient().UpdateTaxRate("c3035aac-27ef-49fd-837a-afebf020b6bc", &UpdateTaxRateDataBody{Name: "MRTT"})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTaxRate(t *testing.T) {
	tst := "TestCreateTaxRate"

	returnValue, err := getClient().CreateTaxRate(&CreateTaxRateDataBody{
		Name: "MRT",
		Rate: "12",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestDeleteTaxRate(t *testing.T) {
	tst := "TestDeleteTaxRate"

	err := getClient().DeleteTaxRate("c3035aac-27ef-49fd-837a-afebf020b6bc")

	OnlyError(t, tst, err)
}
