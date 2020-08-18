package slyk

import "testing"

func TestGetUsers(t *testing.T) {
	tst := "TestGetUsers"

	returnValue, err := getClient().GetUsers()
	ReturnAndError(t, tst, returnValue, err)

	/*
		returnValue, err := getClient().GetUsers(GetUsersFilter().SetEmail("mertacel@gmail.com"))
		ReturnAndError(t, tst, returnValue, err)

			returnValue, err = getClient().GetUsers(GetUsersFilter().
				SetEmail("mertacel@gmail.com").
				SetBlocked(true))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetName("MERT"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetPhone("123"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetPrimaryWalletId("dc80f0c2-3360-46b4-bde0-08ccf06d2878"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetPrimaryWalletId("dc80f0c2-3360-46b4-bde0-08ccf06d2878"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetReferralCode("RV2EXGMMVXG"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetReferralUserID(""))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetRole("owner"))
			ReturnAndError(t, tst, returnValue, err)

			returnValue, err := getClient().GetUsers(GetUsersFilter().SetVerified(true).SetEmail("mertacel@gmail.com"), GetUsersFilter().SetVerified(true).SetReferralCode("RV2EXGMMVXG"))
			ReturnAndError(t, tst, returnValue, err)
	*/
}

func TestGetUserWithID(t *testing.T) {
	tst := "TestGetUserWithID"

	returnValue, err := getClient().GetUserWithID("1bd9b033-5e8a-4d15-bfde-7610f7bd32eb")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateUser(t *testing.T) {
	tst := "TestUpdateUser"

	returnValue, err := getClient().UpdateUser("1bd9b033-5e8a-4d15-bfde-7610f7bd32eb",
		UpdateUserDataForBody().
			SetName("Mert A").
			SetLocale("en").
			SetCustomData(map[string]interface{}{"Test": "Mert"}))

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateUser(t *testing.T) {
	tst := "TestCreateUser"

	returnValue, err := getClient().CreateUser(CreateUserDataForBody().
		SetName("Mert").
		SetEmail("mertacel@gmail.com").
		SetPassword("12345678.aA").
		SetCustomData(map[string]interface{}{"name": "mert"}).
		SetVerified(true).
		SetApproved(true).
		SetBlocked(false).
		SetLocale("en"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestSetUserApprove(t *testing.T) {
	tst := "TestSetUserApprove"

	err := getClient().SetUserApprove("1bd9b033-5e8a-4d15-bfde-7610f7bd32eb")

	OnlyError(t, tst, err)
}

func TestSetUserBlock(t *testing.T) {
	tst := "TestSetUserBlock"

	err := getClient().SetUserBlock("ff17a840-44de-4fdd-94aa-d7619c0aa01b")

	OnlyError(t, tst, err)
}

func TestSetUserUnblock(t *testing.T) {
	tst := "TestSetUserUnblock"

	err := getClient().SetUserUnblock("ff17a840-44de-4fdd-94aa-d7619c0aa01b")

	OnlyError(t, tst, err)
}

func TestChangePassword(t *testing.T) {
	tst := "TestChangePassword"

	err := getClient().ChangePassword("ff17a840-44de-4fdd-94aa-d7619c0aa01b", "12345678.aA")

	OnlyError(t, tst, err)
}
