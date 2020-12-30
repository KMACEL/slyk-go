package slyk

import "time"

type Wallets struct {
	Data  []WalletData `json:"data"`
	Total int          `json:"total"`
}

type Wallet struct {
	Data WalletData `json:"data"`
}

type WalletData struct {
	CreatedAt   time.Time              `json:"createdAt"`
	CustomData  map[string]interface{} `json:"customData"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Locked      bool                   `json:"locked"`
	Metadata    map[string]interface{} `json:"metadata"`
	Name        string                 `json:"name"`
	OwnerID     string                 `json:"ownerId"`
	Reference   string                 `json:"reference"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

type WalletActivities struct {
	Data  []WalletActivityData `json:"data"`
	Total int                  `json:"total"`
}

type WalletActivityData struct {
	Amount             string      `json:"amount"`
	AssetCode          string      `json:"assetCode"`
	Code               string      `json:"code"`
	CreatedAt          time.Time   `json:"createdAt"`
	CustomData         interface{} `json:"customData"`
	DestinationAddress interface{} `json:"destinationAddress,omitempty"`
	DestinationWallet  struct {
		CreatedAt  time.Time   `json:"createdAt"`
		CustomData interface{} `json:"customData"`
		ID         string      `json:"id"`
		Locked     bool        `json:"locked"`
		Metadata   interface{} `json:"metadata"`
		Name       string      `json:"name"`
		OwnerID    string      `json:"ownerId"`
		Reference  string      `json:"reference"`
		UpdatedAt  time.Time   `json:"updatedAt"`
	} `json:"destinationWallet,omitempty"`
	DestinationWalletID   string `json:"destinationWalletId"`
	DestinationWalletUser struct {
		Approved        bool        `json:"approved"`
		Blocked         bool        `json:"blocked"`
		CreatedAt       time.Time   `json:"createdAt"`
		CustomData      interface{} `json:"customData"`
		Email           string      `json:"email"`
		ID              string      `json:"id"`
		Locale          string      `json:"locale"`
		Name            string      `json:"name"`
		Phone           string      `json:"phone"`
		PrimaryWalletID string      `json:"primaryWalletId"`
		ReferralCode    string      `json:"referralCode"`
		Roles           []string    `json:"roles"`
		UpdatedAt       time.Time   `json:"updatedAt"`
		Verified        bool        `json:"verified"`
	} `json:"destinationWalletUser,omitempty"`
	ExternalID     interface{} `json:"externalId"`
	ID             string      `json:"id"`
	Metadata       interface{} `json:"metadata"`
	OriginAddress  interface{} `json:"originAddress,omitempty"`
	OriginWalletID interface{} `json:"originWalletId"`
	Status         string      `json:"status"`
	Type           string      `json:"type"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	OriginWallet   struct {
		CreatedAt  time.Time   `json:"createdAt"`
		CustomData interface{} `json:"customData"`
		ID         string      `json:"id"`
		Locked     bool        `json:"locked"`
		Metadata   interface{} `json:"metadata"`
		Name       string      `json:"name"`
		OwnerID    string      `json:"ownerId"`
		Reference  string      `json:"reference"`
		UpdatedAt  time.Time   `json:"updatedAt"`
	} `json:"originWallet,omitempty"`
	OriginWalletUser struct {
		Approved        bool        `json:"approved"`
		Blocked         bool        `json:"blocked"`
		CreatedAt       time.Time   `json:"createdAt"`
		CustomData      interface{} `json:"customData"`
		Email           string      `json:"email"`
		ID              string      `json:"id"`
		Locale          string      `json:"locale"`
		Name            string      `json:"name"`
		Phone           string      `json:"phone"`
		PrimaryWalletID string      `json:"primaryWalletId"`
		ReferralCode    string      `json:"referralCode"`
		Roles           []string    `json:"roles"`
		UpdatedAt       time.Time   `json:"updatedAt"`
		Verified        bool        `json:"verified"`
	} `json:"originWalletUser,omitempty"`
}

type WalletBalance struct {
	Data []WalletBalanceData `json:"data"`
}

type WalletBalanceData struct {
	AssetCode string `json:"assetCode"`
	Amount    string `json:"amount"`
}

type WalletMovements struct {
	Data  []WalletMovementData `json:"data"`
	Total int                  `json:"total"`
}

type WalletMovementData struct {
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
		CustomData          interface{} `json:"customData"`
		Description         string      `json:"description"`
		DestinationAddress  interface{} `json:"destinationAddress"`
		DestinationWalletID string      `json:"destinationWalletId"`
		ExternalID          interface{} `json:"externalId"`
		ID                  string      `json:"id"`
		Metadata            interface{} `json:"metadata"`
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

type WalletTransactions struct {
	Data  []WalletTransactionData `json:"data"`
	Total int                     `json:"total"`
}

type WalletTransactionData struct {
	Amount              string      `json:"amount"`
	AssetCode           string      `json:"assetCode"`
	Code                string      `json:"code"`
	CreatedAt           time.Time   `json:"createdAt"`
	CustomData          interface{} `json:"customData"`
	Description         interface{} `json:"description"`
	DestinationAddress  interface{} `json:"destinationAddress,omitempty"`
	DestinationWalletID string      `json:"destinationWalletId"`
	ExternalID          interface{} `json:"externalId"`
	ExternalReference   interface{} `json:"externalReference"`
	ID                  string      `json:"id"`
	Metadata            interface{} `json:"metadata"`
	OriginAddress       interface{} `json:"originAddress,omitempty"`
	OriginWalletID      interface{} `json:"originWalletId"`
	ProcessedAt         time.Time   `json:"processedAt"`
	Reference           interface{} `json:"reference"`
	Status              string      `json:"status"`
	Type                string      `json:"type"`
	UpdatedAt           time.Time   `json:"updatedAt"`
}

type UpdateWalletDataBody struct {
	Locked     bool        `json:"locked"`
	OwnerID    string      `json:"ownerId,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}

type CreateWalletDataBody struct {
	Name       string      `json:"name,omitempty"`
	Locked     bool        `json:"locked"`
	OwnerID    string      `json:"ownerId,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}
