package slyk

import "testing"

func TestGetTransactions(t *testing.T) {
	tst := "TestGetTransactions"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetTransactions(GetTransactionsFilter().SetAssetCodeWithIN("usd"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetTransactionsWithID(t *testing.T) {
	tst := "TestGetTransactionsWithID"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetTransactionsWithID("674ae85e-41d9-459c-bdc4-ccf965c19c19")

	ReturnAndError(t, tst, returnValue, err)
}
