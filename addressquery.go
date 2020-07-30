package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetAddresses is
// https://developers.slyk.io/slyk/reference/endpoints#get-addresses
func (c Client) GetAddresses(filter ...*geTransactionstFilter) (*Addresses, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkAddresses)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var addresses Addresses
	errUnmarshal := json.Unmarshal(resp.Body(), &addresses)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &addresses, nil
}

// GetAddressWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-addresses-address
func (c Client) GetAddressWithID(addressID string) (*Address, error) {

	clientReq := resty.New().R()

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkAddresses + ":" + addressID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var address Address
	errUnmarshal := json.Unmarshal(resp.Body(), &address)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &address, nil
}

// CreateAddress is
// https://developers.slyk.io/slyk/reference/endpoints#post-addresses
// TODO : çalışmıyor bakılacak
func (c Client) CreateAddress(createAddressBody *CreateAddressBody) (*Address, error) {

	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(createAddressBody).
		Post(linkAddresses)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var address Address
	errUnmarshal := json.Unmarshal(resp.Body(), &address)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &address, nil
}
