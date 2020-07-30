package slyk

import "time"

type Assests struct {
	Data  []AssetData `json:"data"`
	Total int         `json:"total"`
}

type Asset struct {
	Data AssetData `json:"data"`
}

type AssetData struct {
	Code          string      `json:"code"`
	Contract      struct{}    `json:"contract"`
	CreatedAt     time.Time   `json:"createdAt"`
	CustomData    struct{}    `json:"customData,omitempty"`
	DecimalPlaces int         `json:"decimalPlaces"`
	Enabled       bool        `json:"enabled"`
	Logo          interface{} `json:"logo"`
	Metadata      struct{}    `json:"metadata"`
	Name          string      `json:"name"`
	Symbol        string      `json:"symbol"`
	System        bool        `json:"system"`
	Type          string      `json:"type"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

type UpdateAssetDataBody struct {
	DecimalPlaces int      `json:"decimalPlaces"`
	Name          string   `json:"name"`
	Contract      struct{} `json:"contract"`
	CustomData    struct{} `json:"customData"`
	Enabled       bool     `json:"enabled"`
	Logo          string   `json:"logo"`
	Symbol        string   `json:"symbol"`
}
