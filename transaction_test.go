package slyk

import "testing"

func TestGetTransactions(t *testing.T) {
	tst := "TestGetTransactions"

	//returnValue, err := getClient().GetTransactions()
	/*returnValue, err := getClient().GetTransactions(
	GetTransactionsFilter().
		SetAssetCode("usd").
		SetDestinationWalletIdWithIN("d6f417be-59f8-4728-a782-b810215f0ef2", "73fb8803-bd14-4127-bdb3-8a71b030d4bd"))*/

	returnValue, err := getClient().GetTransactions(GetTransactionsFilter().SetStatus("pending"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetTransactionsWithID(t *testing.T) {
	tst := "TestGetTransactionsWithID"

	returnValue, err := getClient().GetTransactionsWithID("9bd25790-5437-42d3-bab6-da97ed96dcf6")

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTransactionApproveWithID(t *testing.T) {
	tst := "TestCreateTransactionApproveWithID"

	returnValue, err := getClient().SetTransactionApproveWithID("7ce14b27-dfcd-4207-8511-deb2feff0d27")

	ReturnAndError(t, tst, returnValue, err)
}

func TestTransactionConfirmWithID(t *testing.T) {
	tst := "TestTransactionConfirmWithID"

	returnValue, err := getClient().SetTransactionConfirmWithID("e3ca4397-7d4a-498a-b474-2ef58ae313f8")

	ReturnAndError(t, tst, returnValue, err)
}

func TestSetTransactionFailWithID(t *testing.T) {
	tst := "TestSetTransactionFailWithID"

	returnValue, err := getClient().SetTransactionFailWithID("70ec8e59-05c4-4a5e-881b-fa4db1b2d172", &TransactionFailDataBody{Reason: "Problem"})

	ReturnAndError(t, tst, returnValue, err)
}

func TestSetTransactionRejectWithID(t *testing.T) {
	tst := "TestSetTransactionRejectWithID"

	returnValue, err := getClient().SetTransactionRejectWithID("9bd25790-5437-42d3-bab6-da97ed96dcf6", &TransactionRejectDataBody{Reason: "Problem"})

	ReturnAndError(t, tst, returnValue, err)
}

func TestAddTransactionDeposit(t *testing.T) {
	tst := "TestAddTransactionDeposit"

	returnValue, err := getClient().AddTransactionDeposit("327f8888-5f88-4edf-a4b9-7c2066e6e0d9",
		&AddTransactionDepositDataBody{
			Amount:              "0.1",
			AssetCode:           "usd",
			Code:                "internal",
			Data:                map[string]interface{}{"description": "payment"},
			DestinationWalletID: "c553b251-9bf5-42a4-99ef-238ad03fa656",
		})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTransactionPay(t *testing.T) {
	tst := "TestCreateTransactionPay"

	returnValue, err := getClient().CreateTransactionPay(&CreateTransactionPayDataBody{
		Amount:         "12",
		AssetCode:      "usd",
		OriginWalletID: "dc80f0c2-3360-46b4-bde0-08ccf06d2878",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTransactionTransfer(t *testing.T) {
	tst := "TestCreateTransactionTransfer"

	returnValue, err := getClient().CreateTransactionTransfer(&CreateTransactionTransferDataBody{
		Amount:              "1",
		AssetCode:           "usd",
		Code:                "internal",
		DestinationWalletID: "c553b251-9bf5-42a4-99ef-238ad03fa656",
		OriginWalletID:      "d6f417be-59f8-4728-a782-b810215f0ef2",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTransactionWithdrawal(t *testing.T) {
	tst := "TestCreateTransactionWithdrawal"

	returnValue, err := getClient().CreateTransactionWithdrawal(&CreateTransactionWithdrawalDataBody{
		Amount:         "1",
		AssetCode:      "usd",
		Code:           "internal",
		OriginWalletID: "d6f417be-59f8-4728-a782-b810215f0ef2",
		Data:           map[string]interface{}{"Test": "mert"},
	})

	ReturnAndError(t, tst, returnValue, err)
}
