package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c Client) GetTransactions(filter ...*geTransactionstFilter) (*Transaction, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkTransactions)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(resp.Body(), &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}
