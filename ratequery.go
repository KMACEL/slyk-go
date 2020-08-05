package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// CreateRate
// https://developers.slyk.io/slyk/reference/endpoints#post-rates
func (c Client) CreateRate(rateBody *CreateRateBodyData) (*Rate, error) {

	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(rateBody).
		Post(linkRates)

	fmt.Println(string(resp.Body()))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var rate Rate
	errUnmarshal := json.Unmarshal(resp.Body(), &rate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rate, nil
}

// GetRates
// https://developers.slyk.io/slyk/reference/endpoints#get-rates
func (c Client) GetRates(filter ...*getassetFilter) (*Rates, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkRates)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var rates Rates
	errUnmarshal := json.Unmarshal(resp.Body(), &rates)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rates, nil
}

// GetRatesWithBaseAssetCodeAndQuoteAssetCode is
// https://developers.slyk.io/slyk/reference/endpoints#get-rates-baseassetcode-quoteassetcode
func (c Client) GetRatesWithBaseAssetCodeAndQuoteAssetCode(baseAssetCode string, quoteAssetCode string) (*Rate, error) {

	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var rate Rate
	errUnmarshal := json.Unmarshal(resp.Body(), &rate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rate, nil
}

// UpdateRate is
// https://developers.slyk.io/slyk/reference/endpoints#patch-rates-baseassetcode-quoteassetcode
func (c Client) UpdateRate(baseAssetCode string, quoteAssetCode string, rateData string) (*Rate, error) {

	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(&UpdateRateBodyData{Rate: rateData}).
		Patch(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var rate Rate
	errUnmarshal := json.Unmarshal(resp.Body(), &rate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rate, nil
}

// DeleteRate is
// https://developers.slyk.io/slyk/reference/endpoints#delete-rates-baseassetcode-quoteassetcode
func (c Client) DeleteRate(baseAssetCode string, quoteAssetCode string) error {

	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Delete(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	return nil
}
