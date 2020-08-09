package slyk

import "testing"

func TestGetWallet(t *testing.T) {
	tst := "TestGetWallet"

	/*returnValue, err := getClient().GetWallets()

	returnValue, err := getClient().GetWallets(GetWalletFilter().SetIDWithIN("73fb8803-bd14-4127-bdb3-8a71b030d4bd"))
	returnValue, err := getClient().GetWallets(GetWalletFilter().SetIDWithIN("73fb8803-bd14-4127-bdb3-8a71b030d4bd"))
	returnValue, err := getClient().GetWallets(GetWalletFilter().SetLocked(false))
	returnValue, err := getClient().GetWallets(GetWalletFilter().SetName("Master"))
	*/

	returnValue, err := getClient().GetWallets(GetWalletFilter().
		SetIDWithIN("cdbc8ea4-045b-4df2-a1fe-6a286cbdb9c6").
		SetOwnerID("e2cb8c8f-5b87-451e-a0b4-d2bb158b30b2").
		SetName("ACEL2").
		SetLocked(false).
		SetSortWithCreatedAt())

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletWithID(t *testing.T) {
	tst := "TestGetWalletWithID"

	returnValue, err := getClient().GetWalletWithID("73fb8803-bd14-4127-bdb3-8a71b030d4bd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletActivityWithID(t *testing.T) {
	tst := "TestGetWalletActivityWithID"

	returnValue, err := getClient().GetWalletActivityWithID("73fb8803-bd14-4127-bdb3-8a71b030d4bd",
		GetWalletActivtyWithIDFilter().
			SetStatusWithIN("completed").
			SetCodeWithIN("internal"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletBalanceWithID(t *testing.T) {
	tst := "TestGetWalletBalanceWithID"

	returnValue, err := getClient().GetWalletBalanceWithID("d6f417be-59f8-4728-a782-b810215f0ef2",
		GetWalletBalanceWithIDFilter().
			SetAssetCodeWithIN("usd"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletMovements(t *testing.T) {
	tst := "TestGetWalletMovements"

	returnValue, err := getClient().GetWalletMovements("73fb8803-bd14-4127-bdb3-8a71b030d4bd",
		GetWalletMovementFilter().
			SetAssetCodeWithIN("usd").
			SetSortWithAmount())

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletTransactions(t *testing.T) {
	tst := "TestGetWalletTransactions"

	returnValue, err := getClient().GetWalletTransactions("73fb8803-bd14-4127-bdb3-8a71b030d4bd",
		GetWalletTransactionsFilter().
			SetStatusWithIN("completed").
			SetCodeWithIN("internal"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletActivity(t *testing.T) {
	tst := "TestGetWalletActivity"

	returnValue, err := getClient().GetWalletActivity(GetWalletActivityFilter().
		SetAssetCodeWithIN("usd"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletBalance(t *testing.T) {
	tst := "TestGetWalletBalance"

	returnValue, err := getClient().GetWalletBalance()

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateWallet(t *testing.T) {
	tst := "TestUpdateWallet"

	returnValue, err := getClient().UpdateWallet("73fb8803-bd14-4127-bdb3-8a71b030d4bd",
		UpdateWalletBody().
			SetLocked(true).
			SetOwnerID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d").
			SetCustomData(map[string]interface{}{"WalletName": "Test"}))

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateWallet(t *testing.T) {
	tst := "TestCreateWallet"
	returnValue, err := getClient().CreateWallet(CreateWalletBody().
		SetLocked(false).
		SetName("ACEL2").
		SetCustomData(map[string]interface{}{"Name": "Test"}).
		SetOwnerID("e2cb8c8f-5b87-451e-a0b4-d2bb158b30b2"))

	ReturnAndError(t, tst, returnValue, err)
}
