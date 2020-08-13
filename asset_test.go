package slyk

import "testing"

func TestGetAssets(t *testing.T) {
	tst := "TestGetAssets"

	//returnValue, err := getClient().GetAssets()
	returnValue, err := getClient().GetAssets(
		GetAssetFilter().
			SetCodeWithNIN("eur").
			SetPageSize(30))

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

	returnValue, err := getClient().UpdateAssetsWithCode("123456", UpdateAssetDataForBody().SetName("US Dollar"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateAsset(t *testing.T) {
	tst := "TestCreateAsset"

	returnValue, err := getClient().CreateAsset(CreateAssetDataForBody().SetName("Acel").SetCode("ACL").SetType("custom"))

	ReturnAndError(t, tst, returnValue, err)
}
