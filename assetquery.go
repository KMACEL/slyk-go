package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetAssets is
// https://developers.slyk.io/slyk/reference/endpoints#get-assets
func (c Client) GetAssets(filter ...*getassetFilter) (*Assests, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkAssets)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var assests Assests
	errUnmarshal := json.Unmarshal(resp.Body(), &assests)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &assests, nil
}
