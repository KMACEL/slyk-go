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
	Lines                 []struct {
		AssetCode         string                 `json:"assetCode"`
		BonusAssetCode    string                 `json:"bonusAssetCode"`
		CreatedAt         time.Time              `json:"createdAt"`
		FulfilledAt       time.Time              `json:"fulfilledAt"`
		FulfilledQuantity int                    `json:"fulfilledQuantity"`
		Metadata          map[string]interface{} `json:"metadata"`
		ID                string                 `json:"id"`
		OrderID           string                 `json:"orderId"`
		Product           struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"product"`
		ProductID string    `json:"productId"`
		Quantity  int       `json:"quantity"`
		Status    string    `json:"status"`
		UnitBonus string    `json:"unitBonus"`
		UnitPrice string    `json:"unitPrice"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"lines"`
	Metadata       map[string]interface{} `json:"metadata"`
	OrderStatus    string                 `json:"orderStatus"`
	PaidAmount     string                 `json:"paidAmount"`
	PaidAt         time.Time              `json:"paidAt"`
	PaymentStatus  string                 `json:"paymentStatus"`
	ShippingAmount string                 `json:"shippingAmount"`
	TaxesAmount    string                 `json:"taxesAmount"`
	TrackingID     string                 `json:"trackingId"`
	UnpaidAmount   string                 `json:"unpaidAmount"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	User           struct {
		Approved bool                   `json:"approved"`
		Email    string                 `json:"email"`
		ID       string                 `json:"id"`
		ImageURL string                 `json:"imageUrl"`
		Metadata map[string]interface{} `json:"metadata"`
		Name     string                 `json:"name"`
	} `json:"user"`
	UserID string `json:"userId"`
}

type OrderLines struct {
	Data  []OrderLineData `json:"data"`
	Total int             `json:"total"`
}

type OrderLine struct {
	Data OrderLineData `json:"data"`
}

type OrderLineData struct {
	AssetCode         string    `json:"assetCode"`
	BonusAssetCode    string    `json:"bonusAssetCode"`
	CreatedAt         time.Time `json:"createdAt"`
	FulfilledAt       time.Time `json:"fulfilledAt"`
	FulfilledQuantity int       `json:"fulfilledQuantity"`
	ID                string    `json:"id"`
	Metadata          struct{}  `json:"metadata"`
	Order             struct {
		Amount         string      `json:"amount"`
		AssetCode      string      `json:"assetCode"`
		Bonus          string      `json:"bonus"`
		BonusAssetCode string      `json:"bonusAssetCode"`
		CanceledAt     interface{} `json:"canceledAt"`
		CreatedAt      time.Time   `json:"createdAt"`
		CustomData     struct{}    `json:"customData"`
		FulfilledAt    interface{} `json:"fulfilledAt"`
		ID             string      `json:"id"`
		Metadata       struct{}    `json:"metadata"`
		OrderStatus    string      `json:"orderStatus"`
		PaidAmount     string      `json:"paidAmount"`
		PaidAt         interface{} `json:"paidAt"`
		PaymentStatus  string      `json:"paymentStatus"`
		TrackingID     string      `json:"trackingId"`
		UpdatedAt      time.Time   `json:"updatedAt"`
		UserID         string      `json:"userId"`
	} `json:"order"`
	OrderID string `json:"orderId"`
	Product struct {
		Category struct {
			CustomData   struct{}    `json:"customData"`
			Description  string      `json:"description"`
			ID           string      `json:"id"`
			ImageURL     interface{} `json:"imageUrl"`
			Metadata     struct{}    `json:"metadata"`
			ThumbnailURL interface{} `json:"thumbnailUrl"`
			Title        string      `json:"title"`
		} `json:"category"`
		CustomData  struct{}    `json:"customData"`
		Description interface{} `json:"description"`
		ID          string      `json:"id"`
		Image       interface{} `json:"image"`
		Metadata    struct{}    `json:"metadata"`
		Name        string      `json:"name"`
		Thumbnail   interface{} `json:"thumbnail"`
		TypeCode    string      `json:"typeCode"`
	} `json:"product"`
	ProductID string    `json:"productId"`
	Quantity  int       `json:"quantity"`
	Status    string    `json:"status"`
	UnitBonus string    `json:"unitBonus"`
	UnitPrice string    `json:"unitPrice"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateOrderDataBody struct {
	TrackingID string `json:"trackingId"`
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

type OrderPayDataBody struct {
	Amount   string `json:"amount,omitempty"`
	WalletID string `json:"walletId,omitempty"`
}

type OrderLineFulfillDataBody struct {
	Quantity int `json:"quantity"`
}

type OrderLineUNFulfillDataBody struct {
	Quantity int `json:"quantity"`
}
