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

type QuestionTypes struct {
	Data  []QuestionTypeData `json:"data"`
	Total int                `json:"total"`
}

type QuestionTypeData struct {
	Code                string      `json:"code"`
	CreatedAt           time.Time   `json:"createdAt"`
	UpdatedAt           time.Time   `json:"updatedAt"`
	DashboardJSONSchema interface{} `json:"dashboardJsonSchema,omitempty"`
	JSONSchemaTemplate  interface{} `json:"jsonSchemaTemplate,omitempty"`
}

type UpdateQuestionDataBody struct {
	Configurations  interface{} `json:"configurations,omitempty"`
	CustomData      interface{} `json:"customData,omitempty"`
	Description     string      `json:"description,omitempty"`
	ProductTypeCode string      `json:"productTypeCode,omitempty"`
	Required        bool        `json:"required,omitempty"`
	Title           string      `json:"title,omitempty"`
	TypeCode        string      `json:"typeCode,omitempty"`
}
