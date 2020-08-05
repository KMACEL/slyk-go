package slyk

import "time"

type PaymentMedhods struct {
	Data  []PaymentMedhodData `json:"data"`
	Total int                 `json:"total"`
}

type PaymentMedhod struct {
	Data PaymentMedhodData `json:"data"`
}

type PaymentMedhodData struct {
	Assets       []string  `json:"assets"`
	Capabilities []string  `json:"capabilities"`
	CreatedAt    time.Time `json:"createdAt"`
	Enabled      bool      `json:"enabled"`
	Features     []string  `json:"features"`
	Metadata     struct{}  `json:"metadata,omitempty"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
}
