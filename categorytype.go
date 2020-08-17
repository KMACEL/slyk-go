package slyk

import "time"

type Categories struct {
	Data  []CategoryData `json:"data"`
	Total int            `json:"total"`
}

type Category struct {
	Data CategoryData `json:"data"`
}

type CategoryData struct {
	CreatedAt   time.Time              `json:"createdAt"`
	CustomData  map[string]interface{} `json:"customData"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	ImageURL    string                 `json:"imageUrl"`
	Metadata    map[string]interface{} `json:"metadata"`
	Title       string                 `json:"title"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

type UpdateCategoryDataBody struct {
	Description string      `json:"description,omitempty"`
	Image       string      `json:"image,omitempty"`
	Title       string      `json:"title,omitempty"`
	CustomData  interface{} `json:"customData"`
	Order       string      `json:"order,omitempty"`
}

type CreateCategoryDataBody struct {
	Description string      `json:"description,omitempty"`
	Image       string      `json:"image,omitempty"`
	Title       string      `json:"title"`
	CustomData  interface{} `json:"customData"`
	Order       string      `json:"order,omitempty"`
}
