package slyk

import "testing"

func TestGetTasks(t *testing.T) {
	tst := "TestGetTasks"

	returnValue, err := getClient().GetTasks()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetTaskWithID(t *testing.T) {
	tst := "TestGetTaskWithID"

	returnValue, err := getClient().GetTaskWithID("54a5c628-4ab7-435f-8550-39b670da563f")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateTask(t *testing.T) {
	tst := "TestUpdateTask"

	returnValue, err := getClient().UpdateTask("", &UpdateTaskDataBody{})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateTask(t *testing.T) {
	tst := "TestCreateTask"

	returnValue, err := getClient().CreateTask(&CreateTaskDataBody{
		Amount:      "1.0",
		ButtonLabel: "TT",
		Description: "TestTask",
		Name:        "Task",
		Type:        "manual",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestSetTaskComplete(t *testing.T) {
	tst := "TestSetTaskComplete"

	err := getClient().SetTaskComplete("")

	OnlyError(t, tst, err)
}

func TestSetTaskReorder(t *testing.T) {
	tst := "TestSetTaskReorder"

	err := getClient().SetTaskReorder("")

	OnlyError(t, tst, err)
}

func TestDeleteTask(t *testing.T) {
	tst := "TestDeleteTask"

	err := getClient().DeleteTask("")

	OnlyError(t, tst, err)
}
