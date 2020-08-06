package slyk

import "testing"

func TestGetInvites(t *testing.T) {
	tst := "TestGetInvites"

	returnValue, err := getClient().GetInvites()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetInviteWithCode(t *testing.T) {
	tst := "TestGetInviteWithCode"

	returnValue, err := getClient().GetInviteWithCode("IRDE4S3U4Y2")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetInviteWithCodeForValidate(t *testing.T) {
	tst := "TestGetInviteWithCodeForValidate"

	returnValue, err := getClient().GetInviteWithCodeForValidate("IRDE4S3U4Y2")

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateInvite(t *testing.T) {
	tst := "TestCreateInvite"

	returnValue, err := getClient().CreateInvite(&CreateInviteDataBody{
		Email: "mertacel@gmail.com",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCancelInvite(t *testing.T) {
	tst := "TestCancelInvite"

	returnValue, err := getClient().CancelInvite("IRDE4S3U4Y2")

	ReturnAndError(t, tst, returnValue, err)
}

func TestSendInvite(t *testing.T) {
	tst := "TestSendInvite"

	returnValue, err := getClient().SendInvite(&SendInviteDataBody{Email: []string{"mertacel@gmail.com"}, InviterUserID: "cf99e4d8-bc64-4a5c-80a4-dd1e25e2018d"})

	ReturnAndError(t, tst, returnValue, err)
}
