package slyk

import "testing"

func TestGetInvites(t *testing.T) {
	tst := "TestGetInvites"

	returnValue, err := getClient().GetInvites()
	/*returnValue, err := getClient().GetInvites(GetInvitesFilter().
	SetInvitedEmail("mertacel@gmail.com").
	SetCodeWithNIN("I26CFSN63ZP"))*/

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetInviteWithCode(t *testing.T) {
	tst := "TestGetInviteWithCode"

	returnValue, err := getClient().GetInviteWithCode("IDSN91V8S5S")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetInviteWithCodeForValidate(t *testing.T) {
	tst := "TestGetInviteWithCodeForValidate"

	returnValue, err := getClient().GetInviteWithCodeForValidate("IHR7TWS25MA")

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateInvite(t *testing.T) {
	tst := "TestCreateInvite"

	returnValue, err := getClient().CreateInvite(&CreateInviteDataBody{
		Email:         "mertacel@gmail.com",
		InviterUserID: "1bd9b033-5e8a-4d15-bfde-7610f7bd32eb",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCancelInvite(t *testing.T) {
	tst := "TestCancelInvite"

	returnValue, err := getClient().CancelInvite("IHR7TWS25MA")

	ReturnAndError(t, tst, returnValue, err)
}

func TestSendInvite(t *testing.T) {
	tst := "TestSendInvite"

	err := getClient().SendInvite(&SendInviteDataBody{Email: []string{"mertacel@gmail.com"}, InviterUserID: "1bd9b033-5e8a-4d15-bfde-7610f7bd32eb"})

	OnlyError(t, tst, err)
}
