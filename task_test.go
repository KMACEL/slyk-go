package slyk

import "testing"

func TestGetTasks(t *testing.T) {
	tst := "TestGetTasks"

	returnValue, err := getClient().GetTasks()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetTaskWithID(t *testing.T) {
	tst := "TestGetTaskWithID"

	returnValue, err := getClient().GetTaskWithID("")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateTask(t *testing.T) {
	tst := "TestUpdateTask"

	returnValue, err := getClient().UpdateTask("", &UpdateTaskDataBody{})

	ReturnAndError(t, tst, returnValue, err)
}
