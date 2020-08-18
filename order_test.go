package slyk

import "testing"

func TestGetOrders(t *testing.T) {
	tst := "TestGetOrders"

	returnValue, err := getClient().GetOrders()

	ReturnAndError(t, tst, returnValue, err)
}
