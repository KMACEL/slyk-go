package slyk

import "testing"

func TestGetAssets(t *testing.T) {
	tst := "TestGetAssets"

	//returnValue, err := getClient().GetAssets()
	returnValue, err := getClient().GetAssets(GetAssetFilter().SetCodeWithIN("eur"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetAssetsWithCode(t *testing.T) {
	tst := "TestGetAssetsWithCode"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetAssetsWithCode("usd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateAssetsWithCode(t *testing.T) {
	tst := "TestUpdateAssetsWithCode"

	returnValue, err := getClient().UpdateAssetsWithCode("ltc", UpdateAssetDataForBody().SetName("Litecoinn"))

	ReturnAndError(t, tst, returnValue, err)
}
