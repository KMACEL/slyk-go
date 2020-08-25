package slyk

import "testing"

func TestGetTasks(t *testing.T) {
	tst := "TestGetTasks"

	returnValue, err := getClient().GetTasks()

	ReturnAndError(t, tst, returnValue, err)
}
