package slyk

import (
	"encoding/json"
)

// GetWallets is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets
func (c Client) GetWallets(filter ...*getWalletsFilter) (*Wallets, error) {
	getBody, err := c.genericGetQuery(linkWallets, merge(filter))
	if err != nil {
		return nil, err
	}

	var wallets Wallets
	errUnmarshal := json.Unmarshal(getBody, &wallets)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallets, nil
}

// GetWalletWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id
func (c Client) GetWalletWithID(walletID string) (*Wallet, error) {
	getBody, err := c.genericGetQuery(linkWallets+"/"+walletID, nil)
	if err != nil {
		return nil, err
	}

	var wallet Wallet
	errUnmarshal := json.Unmarshal(getBody, &wallet)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallet, nil
}

// GetWalletActivity is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-activity
func (c Client) GetWalletActivityWithID(walletID string, filter ...*getWalletActivityWithIDFilter) (*WalletActivities, error) {
	getBody, err := c.genericGetQuery(linkWallets+"/"+walletID+activity, merge(filter))
	if err != nil {
		return nil, err
	}

	var walletActivities WalletActivities
	errUnmarshal := json.Unmarshal(getBody, &walletActivities)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletActivities, nil
}

// GetWaalletBalance is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-balance
func (c Client) GetWalletBalanceWithID(walletID string, filter ...*getWalletBalanceWithIDFilter) (*WalletBalance, error) {
	getBody, err := c.genericGetQuery(linkWallets+"/"+walletID+balance, merge(filter))
	if err != nil {
		return nil, err
	}

	var walletBalance WalletBalance
	errUnmarshal := json.Unmarshal(getBody, &walletBalance)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletBalance, nil
}

// GetWalletMovements is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-movements
func (c Client) GetWalletMovements(walletID string, filter ...*getWalletMovementsFilter) (*WalletMovements, error) {
	getBody, err := c.genericGetQuery(linkWallets+"/"+walletID+movements, merge(filter))
	if err != nil {
		return nil, err
	}

	var walletMovements WalletMovements
	errUnmarshal := json.Unmarshal(getBody, &walletMovements)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletMovements, nil
}

// GetWalletTransactions is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-transactions
func (c Client) GetWalletTransactions(walletID string, filter ...*getWalletTransactionstFilter) (*WalletTransactions, error) {
	getBody, err := c.genericGetQuery(linkWallets+"/"+walletID+transactions, merge(filter))
	if err != nil {
		return nil, err
	}

	var walletTransactions WalletTransactions
	errUnmarshal := json.Unmarshal(getBody, &walletTransactions)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletTransactions, nil
}

// GetWalletActivity is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-activity
func (c Client) GetWalletActivity(filter ...*getWalletActivityFilter) (*WalletActivities, error) {
	getBody, err := c.genericGetQuery(linkWallets+activity, merge(filter))
	if err != nil {
		return nil, err
	}

	var walletActivities WalletActivities
	errUnmarshal := json.Unmarshal(getBody, &walletActivities)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletActivities, nil
}

// GetWalletBalance is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-balance
func (c Client) GetWalletBalance(filter ...*getWalletBalanceFilter) (*WalletBalance, error) {
	getBody, err := c.genericGetQuery(linkWallets+balance, merge(filter))
	if err != nil {
		return nil, err
	}

	var walletBalance WalletBalance
	errUnmarshal := json.Unmarshal(getBody, &walletBalance)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletBalance, nil
}

// UpdateWallet is
// https://developers.slyk.io/slyk/reference/endpoints#patch-wallets-id
func (c Client) UpdateWallet(walletID string, updateWallet *UpdateWalletDataBody) (*Wallet, error) {
	getBody, err := c.genericPatchQuery(linkWallets+"/"+walletID, updateWallet)
	if err != nil {
		return nil, err
	}

	var wallet Wallet
	errUnmarshal := json.Unmarshal(getBody, &wallet)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallet, nil
}

// CreateWallet is
// https://developers.slyk.io/slyk/reference/endpoints#post-wallets
func (c Client) CreateWallet(createWallet *CreateWalletDataBody) (*Wallet, error) {
	getBody, err := c.genericPostQuery(linkWallets, createWallet)
	if err != nil {
		return nil, err
	}

	var wallet Wallet
	errUnmarshal := json.Unmarshal(getBody, &wallet)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallet, nil
}
