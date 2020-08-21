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

func TestCreateOrder(t *testing.T) {
	tst := "TestCreateOrder"
	//todo producton sonra bakılacak
	returnValue, err := getClient().CreateOrder(&CreateOrderDataBody{
		Lines:  []LineForOrder{{ProductID: ""}},
		UserID: "",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestOrderCancel(t *testing.T) {
	tst := "TestOrderCancel"
	err := getClient().OrderCancel("", &OrderCancelDataBody{})

	OnlyError(t, tst, err)
}

func TestOrderFulfill(t *testing.T) {
	tst := "TestOrderFulfill"
	returnValue, err := getClient().OrderFulfill("", &OrderFulfillDataBody{})

	ReturnAndError(t, tst, returnValue, err)
}

func TestOrderUNFulfill(t *testing.T) {
	tst := "TestOrderUNFulfill"
	returnValue, err := getClient().OrderUNFulfill("")

	ReturnAndError(t, tst, returnValue, err)
}

func TestOrderPay(t *testing.T) {
	tst := "TestOrderPay"
	returnValue, err := getClient().OrderPay("", &OrderPayDataBody{})

	ReturnAndError(t, tst, returnValue, err)
}
