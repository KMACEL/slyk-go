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
func (c Client) GetRates(filter ...*getRateFilter) (*Rates, error) {
	getBody, err := c.genericGetQuery(linkRates, merge(filter))
	if err != nil {
		return nil, err
	}

	var rates Rates
	errUnmarshal := json.Unmarshal(getBody, &rates)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rates, nil
}

// GetRatesWithBaseAssetCodeAndQuoteAssetCode is
// https://developers.slyk.io/slyk/reference/endpoints#get-rates-baseassetcode-quoteassetcode
func (c Client) GetRatesWithBaseAssetCodeAndQuoteAssetCode(baseAssetCode string, quoteAssetCode string) (*Rate, error) {
	getBody, err := c.genericGetQuery(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode), nil)
	if err != nil {
		return nil, err
	}

	var rate Rate
	errUnmarshal := json.Unmarshal(getBody, &rate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rate, nil
}

// UpdateRate is
// https://developers.slyk.io/slyk/reference/endpoints#patch-rates-baseassetcode-quoteassetcode
func (c Client) UpdateRate(baseAssetCode string, quoteAssetCode string, rateData string) (*Rate, error) {
	getBody, err := c.genericPatchQuery(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode), &UpdateRateBodyData{Rate: rateData})
	if err != nil {
		return nil, err
	}

	var rate Rate
	errUnmarshal := json.Unmarshal(getBody, &rate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &rate, nil
}

// DeleteRate is
// https://developers.slyk.io/slyk/reference/endpoints#delete-rates-baseassetcode-quoteassetcode
func (c Client) DeleteRate(baseAssetCode string, quoteAssetCode string) error {
	return c.genericDeleteQuery(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode), nil)
}
