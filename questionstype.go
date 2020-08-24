package slyk

import "time"

type Questions struct {
	Data  []QuestionData `json:"data"`
	Total int            `json:"total"`
}

type Question struct {
	Data QuestionData `json:"data"`
}

type QuestionData struct {
	Configurations  interface{} `json:"configurations"`
	CreatedAt       time.Time   `json:"createdAt"`
	CustomData      interface{} `json:"customData"`
	Description     string      `json:"description"`
	ID              string      `json:"id"`
	JSONSchema      interface{} `json:"jsonSchema"`
	Metadata        interface{} `json:"metadata"`
	ProductTypeCode string      `json:"productTypeCode"`
	Required        bool        `json:"required"`
	Title           string      `json:"title"`
	TypeCode        string      `json:"typeCode"`
	UpdatedAt       time.Time   `json:"updatedAt"`
}
