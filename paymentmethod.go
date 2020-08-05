package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetPaymentMethods is
// https://developers.slyk.io/slyk/reference/endpoints#get-payment-methods
func (c Client) GetPaymentMethods(filter ...*getPaymentMethodFilter) (*PaymentMedhods, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(paymentMethods)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var paymentMedhods PaymentMedhods
	errUnmarshal := json.Unmarshal(resp.Body(), &paymentMedhods)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &paymentMedhods, nil
}

// GetPaymentMethodsWithSlug is
// https://developers.slyk.io/slyk/reference/endpoints#get-payment-methods-slug
func (c Client) GetPaymentMethodsWithSlug(slug string) (*PaymentMedhod, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(paymentMethods + "/" + slug)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var paymentMedhod PaymentMedhod
	errUnmarshal := json.Unmarshal(resp.Body(), &paymentMedhod)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &paymentMedhod, nil
}
