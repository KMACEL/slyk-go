package slyk

import "testing"

func TestGetListFields(t *testing.T) {
	tst := "TestGetListFields"

	returnValue, err := getClient().GetUser()

	ReturnAndError(t, tst, returnValue, err)
}
