package slyk

import "time"

type Transaction struct {
	Data  []TransactionData `json:"data"`
	Total int               `json:"total"`
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
