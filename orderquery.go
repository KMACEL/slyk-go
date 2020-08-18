package slyk

import "encoding/json"

// GetOrders is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-2
func (c Client) GetOrders(filter ...*getOrdersFilter) (*Orders, error) {
	getBody, err := c.genericGetQuery(linkOrders, merge(filter))
	if err != nil {
		return nil, err
	}

	var orders Orders
	errUnmarshal := json.Unmarshal(getBody, &orders)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &orders, nil
}

// GetOrderWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-orders-id
func (c Client) GetOrderWithID(orderID string) (*Order, error) {
	getBody, err := c.genericGetQuery(linkOrders+"/"+orderID, nil)
	if err != nil {
		return nil, err
	}

	var order Order
	errUnmarshal := json.Unmarshal(getBody, &order)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &order, nil
}

// GetOrderLinesWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-3
func (c Client) GetOrderLinesWithID(orderID string, filter ...*getOrderLinesWithIDFilter) (*Orders, error) {
	getBody, err := c.genericGetQuery(linkOrders+"/"+orderID+"/lines", merge(filter))
	if err != nil {
		return nil, err
	}

	var orders Orders
	errUnmarshal := json.Unmarshal(getBody, &orders)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &orders, nil
}

// GetOrderLinesWithIDAndLineID is
// https://developers.slyk.io/slyk/reference/endpoints#get-orders-orderid-lines-id
func (c Client) GetOrderLinesWithIDAndLineID(orderID string, lineID string) (*Order, error) {
	getBody, err := c.genericGetQuery(linkOrders+"/"+orderID+"/lines/"+lineID, nil)
	if err != nil {
		return nil, err
	}

	var order Order
	errUnmarshal := json.Unmarshal(getBody, &order)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &order, nil
}
