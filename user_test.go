package slyk

import "testing"

func TestGetUsers(t *testing.T) {
	tst := "TestGetUsers"

	/*returnValue, err := getClient().GetUser()
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err = getClient().GetUser(CreateUserFilter().SetEmail("mert@monoji.io"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err = getClient().GetUser(CreateUserFilter().
		SetEmail("mert@monoji.io").
		SetBlocked(true))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetName("MERT"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetPhone("123"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetPrimaryWalletId("dc80f0c2-3360-46b4-bde0-08ccf06d2878"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetPrimaryWalletId("dc80f0c2-3360-46b4-bde0-08ccf06d2878"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetReferralCode("RV2EXGMMVXG"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetReferralUserID(""))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(CreateUserFilter().SetRole("owner"))
	ReturnAndError(t, tst, returnValue, err)*/

	returnValue, err := getClient().GetUser(CreateUserFilter().SetVerified(true).SetEmail("mert@monoji.io"), CreateUserFilter().SetVerified(true).SetReferralCode("RV2EXGMMVXG"))
	ReturnAndError(t, tst, returnValue, err)
}

func TestGetUserWithID(t *testing.T) {
	tst := "TestGetUserWithID"

	returnValue, err := getClient().GetUserWithID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d")

	ReturnAndError(t, tst, returnValue, err)
}
