package slyk

import "testing"

func TestGetTransactions(t *testing.T) {
	tst := "TestGetTransactions"

	//returnValue, err := getClient().GetTransactions()
	returnValue, err := getClient().GetTransactions(
		GetTransactionsFilter().
			SetAssetCode("usd").
			SetDestinationWalletIdWithIN("d6f417be-59f8-4728-a782-b810215f0ef2", "73fb8803-bd14-4127-bdb3-8a71b030d4bd"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetTransactionsWithID(t *testing.T) {
	tst := "TestGetTransactionsWithID"

	returnValue, err := getClient().GetTransactionsWithID("674ae85e-41d9-459c-bdc4-ccf965c19c19")

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTransactionApproveWithID(t *testing.T) {
	tst := "TestCreateTransactionApproveWithID"

	returnValue, err := getClient().SetTransactionApproveWithID("674ae85e-41d9-459c-bdc4-ccf965c19c19")

	ReturnAndError(t, tst, returnValue, err)
}

func TestTransactionConfirmWithID(t *testing.T) {
	tst := "TestTransactionConfirmWithID"

	returnValue, err := getClient().SetTransactionConfirmWithID("674ae85e-41d9-459c-bdc4-ccf965c19c19")

	ReturnAndError(t, tst, returnValue, err)
}

func TestSetTransactionFailWithID(t *testing.T) {
	tst := "TestSetTransactionFailWithID"

	returnValue, err := getClient().SetTransactionFailWithID("674ae85e-41d9-459c-bdc4-ccf965c19c19", &TransactionFailDataBody{})

	ReturnAndError(t, tst, returnValue, err)
}

func TestSetTransactionRejectWithID(t *testing.T) {
	tst := "TestSetTransactionRejectWithID"

	returnValue, err := getClient().SetTransactionRejectWithID("674ae85e-41d9-459c-bdc4-ccf965c19c19", &TransactionRejectDataBody{})

	ReturnAndError(t, tst, returnValue, err)
}
