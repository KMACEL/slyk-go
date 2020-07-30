package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetTransactions is
// https://developers.slyk.io/slyk/reference/endpoints#get-transactions
func (c Client) GetTransactions(filter ...*geTransactionstFilter) (*Transactions, error) {

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

	var transactions Transactions
	errUnmarshal := json.Unmarshal(resp.Body(), &transactions)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transactions, nil
}

// GetTransactionsWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-transactions-id
func (c Client) GetTransactionsWithID(transactionID string) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkTransactions + "/" + transactionID)

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