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

// CreateTransactionApproveWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-approve
func (c Client) SetTransactionApproveWithID(transactionID string) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Post(linkTransactions + "/" + transactionID + "/approve")

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

// CreateTransactionConfirmWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-confirm
func (c Client) SetTransactionConfirmWithID(transactionID string) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Post(linkTransactions + "/" + transactionID + "/confirm")

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

// SetTransactionFailWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-fail
func (c Client) SetTransactionFailWithID(transactionID string, transactionFailDataBody *TransactionFailDataBody) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(transactionFailDataBody).
		Post(linkTransactions + "/" + transactionID + "/fail")

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

// SetTransactionRejectWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-reject
func (c Client) SetTransactionRejectWithID(transactionID string, transactionRejectDataBody *TransactionRejectDataBody) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(transactionRejectDataBody).
		Post(linkTransactions + "/" + transactionID + "/reject")

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

// AddTransactionDeposit is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-deposit is
func (c Client) AddTransactionDeposit(transactionID string, addTransactionDepositDataBody *AddTransactionDepositDataBody) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(addTransactionDepositDataBody).
		Post(linkTransactions + "/deposit")

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

// CreateTransaction is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-pay
func (c Client) CreateTransactionPay(createTransactionPayDataBody *CreateTransactionPayDataBody) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(createTransactionPayDataBody).
		Post(linkTransactions + "/pay")

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

// CreateTransactionTransfer
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-transfer
func (c Client) CreateTransactionTransfer(createTransactionTransferDataBody *CreateTransactionTransferDataBody) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(createTransactionTransferDataBody).
		Post(linkTransactions + "/transfer")

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

// CreateTransactionWithdrawal is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-withdrawal
func (c Client) CreateTransactionWithdrawal(createTransactionWithdrawalDataBody *CreateTransactionWithdrawalDataBody) (*Transaction, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(createTransactionWithdrawalDataBody).
		Post(linkTransactions + "/withdrawal")

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
