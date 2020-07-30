package slyk

import "testing"

func TestGetAssets(t *testing.T) {
	tst := "TestGetAssets"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetAssets(GetAssetFilter().SetCodeWithIN("eur"))

	ReturnAndError(t, tst, returnValue, err)
}
