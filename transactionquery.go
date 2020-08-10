package slyk

import (
	"encoding/json"
)

// GetTransactions is
// https://developers.slyk.io/slyk/reference/endpoints#get-transactions
func (c Client) GetTransactions(filter ...*geTransactionstFilter) (*Transactions, error) {
	getBody, err := c.genericGetQuery(linkTransactions, merge(filter))
	if err != nil {
		return nil, err
	}

	var transactions Transactions
	errUnmarshal := json.Unmarshal(getBody, &transactions)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transactions, nil
}

// GetTransactionsWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-transactions-id
func (c Client) GetTransactionsWithID(transactionID string) (*Transaction, error) {
	getBody, err := c.genericGetQuery(linkTransactions+"/"+transactionID, nil)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// CreateTransactionApproveWithID is
// Its only possible to approve transactions that are pending.
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-approve
func (c Client) SetTransactionApproveWithID(transactionID string) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/"+transactionID+"/approve", nil)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// CreateTransactionConfirmWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-confirm
func (c Client) SetTransactionConfirmWithID(transactionID string) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/"+transactionID+"/confirm", nil)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// SetTransactionFailWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-fail
func (c Client) SetTransactionFailWithID(transactionID string, transactionFailDataBody *TransactionFailDataBody) (*Transaction, error) {
	// todo transactionFailDataBody çalışmıyor
	getBody, err := c.genericPostQuery(linkTransactions+"/"+transactionID+"/fail", transactionFailDataBody)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// SetTransactionRejectWithID is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-id-reject
func (c Client) SetTransactionRejectWithID(transactionID string, transactionRejectDataBody *TransactionRejectDataBody) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/"+transactionID+"/reject", transactionRejectDataBody)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// AddTransactionDeposit is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-deposit is
func (c Client) AddTransactionDeposit(transactionID string, addTransactionDepositDataBody *AddTransactionDepositDataBody) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/deposit", addTransactionDepositDataBody)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// CreateTransaction is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-pay
func (c Client) CreateTransactionPay(createTransactionPayDataBody *CreateTransactionPayDataBody) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/pay", createTransactionPayDataBody)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// CreateTransactionTransfer
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-transfer
func (c Client) CreateTransactionTransfer(createTransactionTransferDataBody *CreateTransactionTransferDataBody) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/transfer", createTransactionTransferDataBody)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}

// CreateTransactionWithdrawal is
// https://developers.slyk.io/slyk/reference/endpoints#post-transactions-withdrawal
func (c Client) CreateTransactionWithdrawal(createTransactionWithdrawalDataBody *CreateTransactionWithdrawalDataBody) (*Transaction, error) {
	getBody, err := c.genericPostQuery(linkTransactions+"/withdrawal", createTransactionWithdrawalDataBody)
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	errUnmarshal := json.Unmarshal(getBody, &transaction)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &transaction, nil
}
