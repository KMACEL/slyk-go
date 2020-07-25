package slyk

import "testing"

func TestGetWallet(t *testing.T) {
	tst := "TestGetWallet"

	/*returnValue, err := getClient().GetWallets()

	returnValue, err := getClient().GetWallets(GetWalletFilter().SetIDForIN("73fb8803-bd14-4127-bdb3-8a71b030d4bd"))
	returnValue, err := getClient().GetWallets(GetWalletFilter().SetIDForNsIN("73fb8803-bd14-4127-bdb3-8a71b030d4bd"))
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
