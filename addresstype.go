package slyk

import "time"

type Addresses struct {
	Data  []AddressData `json:"data"`
	Total int           `json:"total"`
}

type Address struct {
	Data  AddressData `json:"data"`
	Total int         `json:"total"`
}

type AddressData struct {
	Address    string    `json:"address"`
	AssetCode  string    `json:"assetCode"`
	CreatedAt  time.Time `json:"createdAt"`
	CustomData struct{}  `json:"customData"`
	Metadata   struct{}  `json:"metadata"`
	UpdatedAt  time.Time `json:"updatedAt"`
	WalletID   string    `json:"walletId"`
}

type CreateAddressBody struct {
	Address    string   `json:"address,omitempty"`
	AssetCode  string   `json:"assetCode,omitempty"`
	CustomData struct{} `json:"customData,omitempty"`
	Provider   string   `json:"provider,omitempty"`
	WalletID   string   `json:"walletId,omitempty"`
}
