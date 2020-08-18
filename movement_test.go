package slyk

import "testing"

func TestGetMovements(t *testing.T) {
	tst := "TestGetMovements"

	returnValue, err := getClient().GetMovements(
		GetMovementsFilter().
			SetAssetCodeWithIN("usd"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetMovementWithID(t *testing.T) {
	tst := "TestGetMovementWithID"

	returnValue, err := getClient().GetMovementWithID("5657aa01-1289-4b38-b7f4-c042423fe053")

	ReturnAndError(t, tst, returnValue, err)
}
