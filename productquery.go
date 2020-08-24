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

// UpdateProducts is
// https://developers.slyk.io/slyk/reference/endpoints#patch-products-id
func (c Client) UpdateProduct(productID string, updateProductsDataBody *UpdateProductDataBody) (*Product, error) {
	getBody, err := c.genericPatchQuery(linkProducts+"/"+productID, updateProductsDataBody)
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
func (c Client) CreateProduct(createProductsDataBody *CreateProductDataBody) (*Product, error) {
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

// AddProductQuestionsWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-products-id-questions
func (c Client) AddProductQuestionWithID(productID string, addProductQuestionDataBody *AddProductQuestionDataBody) (*AddProductQuestionResponseBody, error) {
	getBody, err := c.genericPostQuery(linkProducts+"/"+productID+"/questions", addProductQuestionDataBody)
	if err != nil {
		return nil, err
	}

	var addProductQuestionResponseBody AddProductQuestionResponseBody
	errUnmarshal := json.Unmarshal(getBody, &addProductQuestionResponseBody)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &addProductQuestionResponseBody, nil
}

// AddProductQuestionWithIDForReorder is
// https://developers.slyk.io/slyk/reference/endpoints#post-products-id-reorder
func (c Client) ProductReorder(productID string, productReorderDataBody *ProductReorderDataBody) error {
	_, err := c.genericPostQuery(linkProducts+"/"+productID+"/reorder", productReorderDataBody)
	return err
}

// ProductQuestionReorder is
// https://developers.slyk.io/slyk/reference/endpoints#post-products-productid-questions-id-reorder
func (c Client) ProductQuestionReorder(productID string, questionID string, productQuestionReorderDataBody *ProductQuestionReorderDataBody) error {
	_, err := c.genericPostQuery(linkProducts+"/"+productID+"/questions"+questionID+"/reorder", productQuestionReorderDataBody)
	return err
}

// DeleteProduct is
// https://developers.slyk.io/slyk/reference/endpoints#delete-products-id
func (c Client) DeleteProduct(productID string) error {
	return c.genericDeleteQuery(linkProducts+"/"+productID, nil)
}
