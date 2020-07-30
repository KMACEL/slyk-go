package slyk

import "testing"

func TestGetAssets(t *testing.T) {
	tst := "TestGetAssets"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetAssets(GetAssetFilter().SetCodeWithIN("eur"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetAssetsWithCode(t *testing.T) {
	tst := "TestGetAssetsWithCode"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetAssetsWithCode("ltc")

	ReturnAndError(t, tst, returnValue, err)
}
