package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetWallets is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets
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

// GetWalletWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id
func (c Client) GetWalletWithID(walletID string) (*Wallet, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + "/" + walletID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var wallet Wallet
	errUnmarshal := json.Unmarshal(resp.Body(), &wallet)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallet, nil
}

// GetWalletActivity is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-activity
func (c Client) GetWalletActivityWithID(walletID string, filter ...*getWalletActivityFilter) (*WalletActivities, error) {
	clientReq := resty.New().R()

	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + "/" + walletID + activity)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var walletActivities WalletActivities
	errUnmarshal := json.Unmarshal(resp.Body(), &walletActivities)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletActivities, nil
}

// GetWaalletBalance is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-balance
func (c Client) GetWalletBalanceWithID(walletID string, filter ...*getWalletBalanceWithIDFilter) (*WalletBalance, error) {
	clientReq := resty.New().R()

	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + "/" + walletID + balance)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var walletBalance WalletBalance
	errUnmarshal := json.Unmarshal(resp.Body(), &walletBalance)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletBalance, nil
}

// GetWalletMovements is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-movements
func (c Client) GetWalletMovements(walletID string, filter ...*getWalletMovementFilter) (*WalletMovements, error) {
	clientReq := resty.New().R()

	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + "/" + walletID + movements)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var walletMovements WalletMovements
	errUnmarshal := json.Unmarshal(resp.Body(), &walletMovements)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletMovements, nil
}

// GetWalletTransactions is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-id-transactions
func (c Client) GetWalletTransactions(walletID string, filter ...*getWalletTransactionstFilter) (*WalletTransactions, error) {
	clientReq := resty.New().R()

	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + "/" + walletID + transactions)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var walletTransactions WalletTransactions
	errUnmarshal := json.Unmarshal(resp.Body(), &walletTransactions)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletTransactions, nil
}

// GetWalletActivity is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-activity
func (c Client) GetWalletActivity(filter ...*getWalletActivityFilter) (*WalletActivities, error) {
	clientReq := resty.New().R()

	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + activity)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var walletActivities WalletActivities
	errUnmarshal := json.Unmarshal(resp.Body(), &walletActivities)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletActivities, nil
}

// GetWalletBalance is
// https://developers.slyk.io/slyk/reference/endpoints#get-wallets-balance
func (c Client) GetWalletBalance(filter ...*getWalletBalanceFilter) (*WalletBalance, error) {
	clientReq := resty.New().R()

	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkWallets + balance)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var walletBalance WalletBalance
	errUnmarshal := json.Unmarshal(resp.Body(), &walletBalance)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &walletBalance, nil
}

// UpdateWallet is
// https://developers.slyk.io/slyk/reference/endpoints#patch-wallets-id
// TODO : Çalışmıyror test edilecek
func (c Client) UpdateWallet(walletID string, updateWallet *UpdateWalletData) (*Wallet, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(*updateWallet).
		Patch(linkWallets + "/" + walletID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var wallet Wallet
	errUnmarshal := json.Unmarshal(resp.Body(), &wallet)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallet, nil
}

// CreateWallet is
// https://developers.slyk.io/slyk/reference/endpoints#post-wallets
// TODO Body ile çalışmıyor bakıalcak
func (c Client) CreateWallet(createWallet *CreateWalletData) (*Wallet, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		//SetBody(*createWallet).
		Post(linkWallets)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var wallet Wallet
	errUnmarshal := json.Unmarshal(resp.Body(), &wallet)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &wallet, nil
}
