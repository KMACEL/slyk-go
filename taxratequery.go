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

// GetTaxRateWithID is
//https://developers.slyk.io/slyk/reference/endpoints#get-tax-rates-id
func (c Client) GetTaxRateWithID(taxRateID string) (*TaxRate, error) {
	getBody, err := c.GenericGetQuery(linkTaxRates+"/"+taxRateID, nil)
	if err != nil {
		return nil, err
	}

	var taxRate TaxRate
	errUnmarshal := json.Unmarshal(getBody, &taxRate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &taxRate, nil
}

// UpdateTaxRate is
// https://developers.slyk.io/slyk/reference/endpoints#patch-tax-rates-id
func (c Client) UpdateTaxRate(taxRateID string, updateTaxRateDataBody *UpdateTaxRateDataBody) (*TaxRate, error) {
	getBody, err := c.GenericPatchQuery(linkTaxRates+"/"+taxRateID, updateTaxRateDataBody)
	if err != nil {
		return nil, err
	}

	var taxRate TaxRate
	errUnmarshal := json.Unmarshal(getBody, &taxRate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &taxRate, nil
}
