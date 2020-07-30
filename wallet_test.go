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

	returnValue, err := getClient().GetWallets(GetWalletFilter().SetSortWithCreatedAt())

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletWithID(t *testing.T) {
	tst := "TestGetWalletWithID"

	returnValue, err := getClient().GetWalletWithID("73fb8803-bd14-4127-bdb3-8a71b030d4bd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletActivityWithID(t *testing.T) {
	tst := "TestGetWalletActivityWithID"

	returnValue, err := getClient().GetWalletActivityWithID("73fb8803-bd14-4127-bdb3-8a71b030d4bd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletBalanceWithID(t *testing.T) {
	tst := "TestGetWalletBalanceWithID"

	returnValue, err := getClient().GetWalletBalanceWithID("73fb8803-bd14-4127-bdb3-8a71b030d4bd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletMovements(t *testing.T) {
	tst := "TestGetWalletMovements"

	returnValue, err := getClient().GetWalletMovements("73fb8803-bd14-4127-bdb3-8a71b030d4bd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletTransactions(t *testing.T) {
	tst := "TestGetWalletTransactions"

	returnValue, err := getClient().GetWalletTransactions("73fb8803-bd14-4127-bdb3-8a71b030d4bd")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletActivity(t *testing.T) {
	tst := "TestGetWalletActivity"

	returnValue, err := getClient().GetWalletActivity(GetWalletActivityFilter().SetAssetCodeWithIN("usd"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetWalletBalance(t *testing.T) {
	tst := "TestGetWalletBalance"

	returnValue, err := getClient().GetWalletBalance()

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateWallet(t *testing.T) {
	tst := "TestUpdateWallet"

	returnValue, err := getClient().UpdateWallet("73fb8803-bd14-4127-bdb3-8a71b030d4bd", UpdateWalletBody().SetLocked(true))

	ReturnAndError(t, tst, returnValue, err)
}
