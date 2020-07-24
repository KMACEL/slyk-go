package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c Client) GetWallets(filter ...*getWalletFilter) (*Wallets, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var wallets Wallets
	errUnmarshal := json.Unmarshal(resp.Body(), &wallets)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallets, nil
}
