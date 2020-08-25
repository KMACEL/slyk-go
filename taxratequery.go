package slyk

import "encoding/json"

// GetTaxRates is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-10
func (c Client) GetTaxRates(filter ...*getTaxRatesFilter) (*TaxRates, error) {
	getBody, err := c.GenericGetQuery(linkTaxRates, merge(filter))
	if err != nil {
		return nil, err
	}

	var taxRates TaxRates
	errUnmarshal := json.Unmarshal(getBody, &taxRates)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &taxRates, nil
}
