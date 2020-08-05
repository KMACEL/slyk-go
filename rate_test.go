package slyk

import "testing"

func TestCreateRate(t *testing.T) {
	tst := "TestCreateRate"

	returnValue, err := getClient().CreateRate(&CreateRateBodyData{
		BaseAssetCode:  "btc",
		QuoteAssetCode: "eth",
		Rate:           "48.86",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetRates(t *testing.T) {
	tst := "TestGetRates"

	returnValue, err := getClient().GetRates()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetRatesWithBaseAssetCodeAndQuoteAssetCode(t *testing.T) {
	tst := "TestGetRatesWithBaseAssetCodeAndQuoteAssetCode"

	returnValue, err := getClient().GetRatesWithBaseAssetCodeAndQuoteAssetCode("btc", "eth")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateRate(t *testing.T) {
	tst := "TestUpdateRate"

	returnValue, err := getClient().UpdateRate("btc", "eth", "49")

	ReturnAndError(t, tst, returnValue, err)
}

func TestDeleteRate(t *testing.T) {
	tst := "TestDeleteRate"

	err := getClient().DeleteRate("btc", "eth")

	OnlyError(t, tst, err)
}
