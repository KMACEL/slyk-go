package slyk

import "testing"

func TestGetQuestions(t *testing.T) {
	tst := "TestGetQuestions"

	returnValue, err := getClient().GetQuestions()

	ReturnAndError(t, tst, returnValue, err)
}
