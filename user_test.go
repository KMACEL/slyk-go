package slyk

import "testing"

func TestGetUsers(t *testing.T) {
	tst := "TestGetUsers"

	returnValue, err := getClient().GetUser()
	ReturnAndError(t, tst, returnValue, err)

	/*returnValue, err = getClient().GetUser(GetUserFilter().SetEmail("mert@monoji.io"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err = getClient().GetUser(GetUserFilter().
		SetEmail("mert@monoji.io").
		SetBlocked(true))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetName("MERT"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetPhone("123"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetPrimaryWalletId("dc80f0c2-3360-46b4-bde0-08ccf06d2878"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetPrimaryWalletId("dc80f0c2-3360-46b4-bde0-08ccf06d2878"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetReferralCode("RV2EXGMMVXG"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetReferralUserID(""))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetRole("owner"))
	ReturnAndError(t, tst, returnValue, err)

	returnValue, err := getClient().GetUser(GetUserFilter().SetVerified(true).SetEmail("mert@monoji.io"), GetUserFilter().SetVerified(true).SetReferralCode("RV2EXGMMVXG"))
	ReturnAndError(t, tst, returnValue, err)*/
}

func TestGetUserWithID(t *testing.T) {
	tst := "TestGetUserWithID"

	returnValue, err := getClient().GetUserWithID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d")

	ReturnAndError(t, tst, returnValue, err)
}

func TestPatchUser(t *testing.T) {
	tst := "TestPatchUser"

	returnValue, err := getClient().PatchUser("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d", CreatePatchData().SetName("Mert"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateUser(t *testing.T) {
	tst := "TestCreateUser"

	returnValue, err := getClient().PostCreateUser(CreatePostUser().
		SetName("Mert").
		SetEmail("mertacel@gmail.com").
		SetPassword("123456789..aA"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestPostUserApprove(t *testing.T) {
	tst := "TestPostUserApprove"

	err := getClient().PostUserApprove("debf8f0c-548b-4a35-833a-a5f33dc154ae")

	OnlyError(t, tst, err)
}

func TestPostUserBlock(t *testing.T) {
	tst := "TestPostUserBlock"

	err := getClient().PostUserBlock("debf8f0c-548b-4a35-833a-a5f33dc154ae")

	OnlyError(t, tst, err)
}

func TestPostUserUnblock(t *testing.T) {
	tst := "TestPostUserUnblock"

	err := getClient().PostUserUnblock("debf8f0c-548b-4a35-833a-a5f33dc154ae")

	OnlyError(t, tst, err)
}
