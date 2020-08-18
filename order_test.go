package slyk

import "testing"

func TestGetOrders(t *testing.T) {
	tst := "TestGetOrders"

	returnValue, err := getClient().GetOrders()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetOrderWithID(t *testing.T) {
	tst := "TestGetOrderWithID"

	returnValue, err := getClient().GetOrderWithID("")

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetOrderLinesWithID(t *testing.T) {
	tst := "TestGetOrderLinesWithID"

	returnValue, err := getClient().GetOrderLinesWithID("")

	ReturnAndError(t, tst, returnValue, err)
}
