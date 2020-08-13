package slyk

import "time"

type Movements struct {
	Data  []MovementData `json:"data"`
	Total int            `json:"total"`
}

type Movement struct {
	Data MovementData `json:"data"`
}

type MovementData struct {
	Amount        string    `json:"amount"`
	AssetCode     string    `json:"assetCode"`
	Code          string    `json:"code"`
	CreatedAt     time.Time `json:"createdAt"`
	ID            string    `json:"id"`
	TransactionID string    `json:"transactionId"`
	UpdatedAt     time.Time `json:"updatedAt"`
	WalletID      string    `json:"walletId"`
}
