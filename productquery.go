package slyk

import "encoding/json"

// GetProducts is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-6
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

// GetProductsWithID
// https://developers.slyk.io/slyk/reference/endpoints#get-products-id
func (c Client) GetProductWithID(productID string) (*Product, error) {
	getBody, err := c.genericGetQuery(linkProducts+"/"+productID, nil)
	if err != nil {
		return nil, err
	}

	var product Product
	errUnmarshal := json.Unmarshal(getBody, &product)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &product, nil
}

// CreateProducts is
// https://developers.slyk.io/slyk/reference/endpoints#post-products
func (c Client) CreateProducts(createProductsDataBody *CreateProductsDataBody) (*Product, error) {
	getBody, err := c.genericPostQuery(linkProducts, createProductsDataBody)
	if err != nil {
		return nil, err
	}

	var product Product
	errUnmarshal := json.Unmarshal(getBody, &product)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &product, nil
}
