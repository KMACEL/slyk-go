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

	returnValue, err := getClient().CreateAddress(
		CreateAddressForBody().
			SetAssetCode("usd").
			SetProvider("coinbase").
			SetWalletID("d6f417be-59f8-4728-a782-b810215f0ef2"))

	ReturnAndError(t, tst, returnValue, err)
}
