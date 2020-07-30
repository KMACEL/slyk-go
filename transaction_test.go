package slyk

import "testing"

func TestGetTransactions(t *testing.T) {
	tst := "TestGetTransactions"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetTransactions(GetTransactionsFilter().SetAssetCodeWithIN("usd"))

	ReturnAndError(t, tst, returnValue, err)
}
