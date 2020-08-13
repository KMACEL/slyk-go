package slyk

import "testing"

func TestGetPaymentMethods(t *testing.T) {
	tst := "TestGetPaymentMethods"

	returnValue, err := getClient().GetPaymentMethods(
		GetPaymentMedhodFilter().
			SetEnabled(false))

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetPaymentMethodsWithSlug(t *testing.T) {
	tst := "TestGetPaymentMethodsWithSlug"

	returnValue, err := getClient().GetPaymentMethodsWithSlug("stripe")

	ReturnAndError(t, tst, returnValue, err)
}
