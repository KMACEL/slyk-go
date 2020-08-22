package slyk

import "time"

type Products struct {
	Data  []ProductData `json:"data"`
	Total int           `json:"total"`
}

type Product struct {
	Data ProductData `json:"data"`
}

type ProductData struct {
	AllowChoosingQuantity bool      `json:"allowChoosingQuantity"`
	AssetCode             string    `json:"assetCode"`
	Available             bool      `json:"available"`
	Bonus                 string    `json:"bonus"`
	ButtonLabel           string    `json:"buttonLabel"`
	CategoryID            string    `json:"categoryId"`
	CreatedAt             time.Time `json:"createdAt"`
	CustomData            struct {
	} `json:"customData"`
	Description string `json:"description"`
	Featured    bool   `json:"featured"`
	ID          string `json:"id"`
	ImageURL    string `json:"imageUrl"`
	ListLabel   string `json:"listLabel"`
	Metadata    struct {
	} `json:"metadata"`
	Name             string    `json:"name"`
	Price            string    `json:"price"`
	PriceWithTax     string    `json:"priceWithTax"`
	RequiresIdentity bool      `json:"requiresIdentity"`
	TaxAmount        string    `json:"taxAmount"`
	TaxRateID        string    `json:"taxRateId"`
	ThumbnailURL     string    `json:"thumbnailUrl"`
	UpdatedAt        time.Time `json:"updatedAt"`
	URL              string    `json:"url"`
	Visible          bool      `json:"visible"`
}
