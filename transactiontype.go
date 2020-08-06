package slyk

import "time"

type Transactions struct {
	Data  []TransactionData `json:"data"`
	Total int               `json:"total"`
}

type Transaction struct {
	Data  TransactionData `json:"data"`
	Total int             `json:"total"`
}

type TransactionData struct {
	Amount      string    `json:"amount"`
	AssetCode   string    `json:"assetCode"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	ID          string    `json:"id"`
	Transaction struct {
		Amount              string      `json:"amount"`
		AssetCode           string      `json:"assetCode"`
		Code                string      `json:"code"`
		CreatedAt           time.Time   `json:"createdAt"`
		CustomData          struct{}    `json:"customData"`
		Description         string      `json:"description"`
		DestinationAddress  interface{} `json:"destinationAddress"`
		DestinationWalletID string      `json:"destinationWalletId"`
		ExternalID          interface{} `json:"externalId"`
		ID                  string      `json:"id"`
		Metadata            struct{}    `json:"metadata"`
		OriginAddress       interface{} `json:"originAddress"`
		OriginWalletID      interface{} `json:"originWalletId"`
		Status              string      `json:"status"`
		Type                string      `json:"type"`
		UpdatedAt           time.Time   `json:"updatedAt"`
	} `json:"transaction"`
	TransactionID string    `json:"transactionId"`
	UpdatedAt     time.Time `json:"updatedAt"`
	WalletID      string    `json:"walletId"`
}

type TransactionFailDataBody struct {
	Reason string `json:"reason"`
}

type TransactionRejectDataBody struct {
	Reason string `json:"reason"`
}

type AddTransactionDepositDataBody struct {
	Amount     string   `json:"amount"`
	AssetCode  string   `json:"assetCode"`
	Code       string   `json:"code"`
	CustomData struct{} `json:"customData,omitempty"`
	Data       struct {
		Description string `json:"description"`
	} `json:"data"`
	Description         string `json:"description,omitempty"`
	DestinationAddress  string `json:"destinationAddress,omitempty"`
	DestinationWalletID string `json:"destinationWalletId"`
	Commit              bool   `json:"commit,omitempty"`
	ExternalReference   string `json:"externalReference,omitempty"`
}

type CreateTransactionPayDataBody struct {
	Amount         string   `json:"amount"`
	AssetCode      string   `json:"assetCode"`
	CustomData     struct{} `json:"customData,omitempty"`
	Description    string   `json:"description,omitempty"`
	OriginWalletID string   `json:"originWalletId"`
}

type CreateTransactionTransferDataBody struct {
	Amount              string   `json:"amount"`
	AssetCode           string   `json:"assetCode"`
	Code                string   `json:"code"`
	Commit              bool     `json:"commit"`
	CustomData          struct{} `json:"customData"`
	Description         string   `json:"description"`
	DestinationAddress  string   `json:"destinationAddress"`
	DestinationWalletID string   `json:"destinationWalletId"`
	OriginAddress       string   `json:"originAddress"`
	OriginWalletID      string   `json:"originWalletId"`
}

type CreateTransactionWithdrawalDataBody struct {
	Amount     string   `json:"amount"`
	AssetCode  string   `json:"assetCode"`
	Code       string   `json:"code"`
	CustomData struct{} `json:"customData"`
	Data       struct {
		BankAddress string `json:"bankAddress"`
		BankName    string `json:"bankName"`
		Country     string `json:"country"`
		Iban        string `json:"iban"`
		SwiftCode   string `json:"swiftCode"`
	} `json:"data"`
	Description    string `json:"description"`
	OriginWalletID string `json:"originWalletId"`
}
