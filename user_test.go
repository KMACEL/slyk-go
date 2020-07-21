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
	ReturnAndError(t, tst, returnValue, err)*/

	returnValue, err := getClient().GetUser(CreateUserFilter().SetID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d"))
	ReturnAndError(t, tst, returnValue, err)

}

func TestGetUserWithID(t *testing.T) {
	tst := "TestGetUserWithID"

	returnValue, err := getClient().GetUserWithID("cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d")

	ReturnAndError(t, tst, returnValue, err)
}
