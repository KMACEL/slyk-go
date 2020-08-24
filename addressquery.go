package slyk

import (
	"encoding/json"
)

// GetAddresses is
// https://developers.slyk.io/slyk/reference/endpoints#get-addresses
func (c Client) GetAddresses(filter ...*getaddressesFilter) (*Addresses, error) {
	getBody, err := c.GenericGetQuery(linkAddresses, merge(filter))
	if err != nil {
		return nil, err
	}

	var addresses Addresses
	errUnmarshal := json.Unmarshal(getBody, &addresses)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &addresses, nil
}

// GetAddressWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-addresses-address
func (c Client) GetAddressWithID(addressID string) (*Address, error) {
	getBody, err := c.GenericGetQuery(linkAddresses+"/"+addressID, nil)
	if err != nil {
		return nil, err
	}

	var address Address
	errUnmarshal := json.Unmarshal(getBody, &address)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &address, nil
}

// CreateAddress is
// https://developers.slyk.io/slyk/reference/endpoints#post-addresses
// Only COin Base
func (c Client) CreateAddress(createAddressBody *CreateAddressBody) (*Address, error) {
	getBody, err := c.GenericPostQuery(linkAddresses, createAddressBody)
	if err != nil {
		return nil, err
	}

	var address Address
	errUnmarshal := json.Unmarshal(getBody, &address)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &address, nil
}
