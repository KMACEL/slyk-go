package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetMovements
// https://developers.slyk.io/slyk/reference/endpoints#get-movements
func (c Client) GetMovements(filter ...*getMovementFilter) (*Movements, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkMovements)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var movements Movements
	errUnmarshal := json.Unmarshal(resp.Body(), &movements)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &movements, nil
}

// GetMovementWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-movements-id
func (c Client) GetMovementWithID(movementID string) (*Movement, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkMovements + "/" + movementID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var movement Movement
	errUnmarshal := json.Unmarshal(resp.Body(), &movement)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &movement, nil
}
