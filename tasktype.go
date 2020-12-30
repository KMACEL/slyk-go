package slyk

import "time"

type Tasks struct {
	Data  []TaskData `json:"data"`
	Total int        `json:"total"`
}

type Task struct {
	Data TaskData `json:"data"`
}

type TaskData struct {
	Amount       string                 `json:"amount"`
	Available    bool                   `json:"available"`
	ButtonLabel  string                 `json:"buttonLabel"`
	CreatedAt    time.Time              `json:"createdAt"`
	CustomData   map[string]interface{} `json:"customData"`
	Description  string                 `json:"description"`
	Enabled      bool                   `json:"enabled"`
	Featured     bool                   `json:"featured"`
	ID           string                 `json:"id"`
	ImageURL     string                 `json:"imageUrl"`
	Metadata     map[string]interface{} `json:"metadata"`
	Name         string                 `json:"name"`
	SurveyURL    string                 `json:"surveyUrl"`
	ThumbnailURL string                 `json:"thumbnailUrl"`
	Type         string                 `json:"type"`
	UpdatedAt    time.Time              `json:"updatedAt"`
}

type UpdateTaskDataBody struct {
	Amount      string      `json:"amount,omitempty"`
	ButtonLabel string      `json:"buttonLabel,omitempty"`
	CustomData  interface{} `json:"customData,omitempty"`
	Description string      `json:"description,omitempty"`
	Enabled     bool        `json:"enabled"`
	Featured    bool        `json:"featured"`
	Image       string      `json:"image,omitempty"`
	Name        string      `json:"name,omitempty"`
	Order       string      `json:"order,omitempty"`
	SurveyURL   string      `json:"surveyUrl,omitempty"`
	Thumbnail   string      `json:"thumbnail,omitempty"`
	Type        string      `json:"type,omitempty"`
}

type CreateTaskDataBody struct {
	Amount      string      `json:"amount"`
	ButtonLabel string      `json:"buttonLabel,omitempty"`
	CustomData  interface{} `json:"customData,omitempty"`
	Description string      `json:"description"`
	Enabled     bool        `json:"enabled"`
	Featured    bool        `json:"featured"`
	Image       string      `json:"image,omitempty"`
	Name        string      `json:"name"`
	Order       string      `json:"order,omitempty"`
	SurveyURL   string      `json:"surveyUrl,omitempty"`
	Thumbnail   string      `json:"thumbnail,omitempty"`
	Type        string      `json:"type"`
}
