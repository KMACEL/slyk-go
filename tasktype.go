package slyk

import "time"

type Tasks struct {
	Data  []TaskData `json:"data"`
	Total int        `json:"total"`
}

type TaskData struct {
	Amount       string      `json:"amount"`
	Available    bool        `json:"available"`
	ButtonLabel  string      `json:"buttonLabel"`
	CreatedAt    time.Time   `json:"createdAt"`
	CustomData   interface{} `json:"customData"`
	Description  string      `json:"description"`
	Enabled      bool        `json:"enabled"`
	Featured     bool        `json:"featured"`
	ID           string      `json:"id"`
	ImageURL     string      `json:"imageUrl"`
	Metadata     interface{} `json:"metadata"`
	Name         string      `json:"name"`
	SurveyURL    string      `json:"surveyUrl"`
	ThumbnailURL string      `json:"thumbnailUrl"`
	Type         string      `json:"type"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}
