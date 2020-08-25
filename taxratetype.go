package slyk

import "time"

type TaxRates struct {
	Data  []TaxRateData `json:"data"`
	Total int           `json:"total"`
}

type TaxRate struct {
	Data TaxRateData `json:"data"`
}

type TaxRateData struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Rate      string    `json:"rate"`
	UpdatedAt time.Time `json:"updatedAt"`
}
