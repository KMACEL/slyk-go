package slyk

import "testing"

func TestGetAddress(t *testing.T) {
	tst := "TestGetAddress("

	returnValue, err := getClient().GetAddresses()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetAddressesWithID(t *testing.T) {
	tst := "TestGetAddressesWithID"

	returnValue, err := getClient().GetAddressWithID("")

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateAddress(t *testing.T) {
	tst := "TestCreateAddress"

	returnValue, err := getClient().CreateAddress(CreateAddressForBody().SetAssetCode("ltc").SetProvider("coinbase").SetWalletID("73fb8803-bd14-4127-bdb3-8a71b030d4bd"))

	ReturnAndError(t, tst, returnValue, err)
}
