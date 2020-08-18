package slyk

import (
	"encoding/json"
	"fmt"
)

// GetRates
// https://developers.slyk.io/slyk/reference/endpoints#get-rates
func (c Client) GetRates(filter ...*getRatesFilter) (*Rates, error) {
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
func (c Client) UpdateRate(baseAssetCode string, quoteAssetCode string, updateRateDataBody *UpdateRateDataBody) (*Rate, error) {
	getBody, err := c.genericPatchQuery(fmt.Sprintf("%s/%s/%s", linkRates, baseAssetCode, quoteAssetCode), updateRateDataBody)
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

// CreateRate
// https://developers.slyk.io/slyk/reference/endpoints#post-rates
func (c Client) CreateRate(rateBody *CreateRateDataBody) (*Rate, error) {
	getBody, err := c.genericPostQuery(linkRates, rateBody)
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
