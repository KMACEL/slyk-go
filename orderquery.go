package slyk

import "encoding/json"

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
