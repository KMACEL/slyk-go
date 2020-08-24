package slyk

import (
	"encoding/json"
)

// GetPaymentMethods is
// https://developers.slyk.io/slyk/reference/endpoints#get-payment-methods
func (c Client) GetPaymentMethods(filter ...*getPaymentMethodsFilter) (*PaymentMedhods, error) {
	getBody, err := c.GenericGetQuery(linkPaymentMethods, merge(filter))
	if err != nil {
		return nil, err
	}

	var paymentMedhods PaymentMedhods
	errUnmarshal := json.Unmarshal(getBody, &paymentMedhods)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &paymentMedhods, nil
}

// GetPaymentMethodsWithSlug is
// https://developers.slyk.io/slyk/reference/endpoints#get-payment-methods-slug
func (c Client) GetPaymentMethodsWithSlug(slug string) (*PaymentMedhod, error) {
	getBody, err := c.GenericGetQuery(linkPaymentMethods+"/"+slug, nil)
	if err != nil {
		return nil, err
	}

	var paymentMedhod PaymentMedhod
	errUnmarshal := json.Unmarshal(getBody, &paymentMedhod)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &paymentMedhod, nil
}
