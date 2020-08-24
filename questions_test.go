package slyk

import "testing"

func TestGetQuestions(t *testing.T) {
	tst := "TestGetQuestions"

	returnValue, err := getClient().GetQuestions()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetQuestionWithID(t *testing.T) {
	tst := "TestGetQuestionsWithID"

	returnValue, err := getClient().GetQuestionWithID("d448e2e5-3b79-4ece-8a0f-754c9b1e0d8c")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetQuestionsTypes(t *testing.T) {
	tst := "TestGetQuestionsTypes"

	returnValue, err := getClient().GetQuestionsTypes()

	ReturnAndError(t, tst, returnValue, err)
}
