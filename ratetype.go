package slyk

import "time"

type Rates struct {
	Data  []RateData `json:"data"`
	Total int        `json:"total"`
}

type Rate struct {
	Data RateData `json:"data"`
}

type RateData struct {
	BaseAssetCode  string      `json:"baseAssetCode"`
	CreatedAt      time.Time   `json:"createdAt"`
	CustomData     interface{} `json:"customData"`
	Metadata       interface{} `json:"metadata"`
	QuoteAssetCode string      `json:"quoteAssetCode"`
	Rate           string      `json:"rate"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

type CreateRateBodyData struct {
	BaseAssetCode  string      `json:"baseAssetCode"`
	QuoteAssetCode string      `json:"quoteAssetCode"`
	Rate           string      `json:"rate"`
	CustomData     interface{} `json:"customData,omitempty"`
}

type UpdateRateBodyData struct {
	Rate       string      `json:"rate"`
	CustomData interface{} `json:"customData,omitempty"`
}
