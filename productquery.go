package slyk

import "encoding/json"

func (c Client) GetProducts(filter ...*getProductsFilter) (*Products, error) {
	getBody, err := c.genericGetQuery(linkProducts, merge(filter))
	if err != nil {
		return nil, err
	}

	var products Products
	errUnmarshal := json.Unmarshal(getBody, &products)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &products, nil
}
