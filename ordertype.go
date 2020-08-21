package slyk

import "time"

type Orders struct {
	Data  []OrderData `json:"data"`
	Total int         `json:"total"`
}

type Order struct {
	Data OrderData `json:"data"`
}

type OrderData struct {
	Amount                string                 `json:"amount"`
	AmountWithoutShipment string                 `json:"amountWithoutShipment"`
	AmountWithoutTaxes    string                 `json:"amountWithoutTaxes"`
	AssetCode             string                 `json:"assetCode"`
	Bonus                 string                 `json:"bonus"`
	BonusAssetCode        string                 `json:"bonusAssetCode"`
	CanceledAt            time.Time              `json:"canceledAt"`
	CreatedAt             time.Time              `json:"createdAt"`
	CustomData            map[string]interface{} `json:"customData"`
	DeliveryMethod        string                 `json:"deliveryMethod"`
	FulfilledAt           time.Time              `json:"fulfilledAt"`
	ID                    string                 `json:"id"`
	Metadata              map[string]interface{} `json:"metadata"`
	OrderStatus           string                 `json:"orderStatus"`
	PaidAmount            string                 `json:"paidAmount"`
	PaidAt                time.Time              `json:"paidAt"`
	PaymentStatus         string                 `json:"paymentStatus"`
	ShippingAmount        string                 `json:"shippingAmount"`
	TaxesAmount           string                 `json:"taxesAmount"`
	TrackingID            string                 `json:"trackingId"`
	UnpaidAmount          string                 `json:"unpaidAmount"`
	UpdatedAt             time.Time              `json:"updatedAt"`
	User                  struct {
		Approved bool                   `json:"approved"`
		Email    string                 `json:"email"`
		ID       string                 `json:"id"`
		ImageURL string                 `json:"imageUrl"`
		Metadata map[string]interface{} `json:"metadata"`
		Name     string                 `json:"name"`
	} `json:"user"`
	UserID string `json:"userId"`
}

type CreateOrderDataBody struct {
	ChosenPaymentMethod string         `json:"chosenPaymentMethod,omitempty"`
	CustomData          interface{}    `json:"customData,omitempty"`
	DeliveryMethod      string         `json:"deliveryMethod,omitempty"`
	DryRun              bool           `json:"dryRun,omitempty"`
	Lines               []LineForOrder `json:"lines"`
	ShippingAddressID   string         `json:"shippingAddressId,omitempty"`
	UseBonus            bool           `json:"useBonus,omitempty"`
	UserID              string         `json:"userId"`
	UserNotes           string         `json:"userNotes,omitempty"`
	WalletID            string         `json:"walletId,omitempty"`
}

type LineForOrder struct {
	ProductID              string      `json:"productId"`
	Quantity               int         `json:"quantity,omitempty"`
	ProductQuestionsResult interface{} `json:"productQuestionsResult,omitempty"`
}

type OrderCancelDataBody struct {
	Reason       string `json:"reason,omitempty"`
	RefundAmount string `json:"refundAmount,omitempty"`
}

type OrderFulfillDataBody struct {
	TrackingID string `json:"trackingId"`
}
