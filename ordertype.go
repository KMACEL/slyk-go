package slyk

import "time"

type Orders struct {
	Data  []OrderData `json:"data"`
	Total int         `json:"total"`
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
