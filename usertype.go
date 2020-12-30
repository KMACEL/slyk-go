package slyk

import "time"

type Users struct {
	Data  []UserData `json:"data"`
	Total int        `json:"total"`
}

type User struct {
	UserData `json:"data"`
}

type UserData struct {
	Approved        bool                   `json:"approved"`
	Blocked         bool                   `json:"blocked"`
	CreatedAt       time.Time              `json:"createdAt"`
	CustomData      map[string]interface{} `json:"customData"`
	Email           string                 `json:"email"`
	ID              string                 `json:"id"`
	Image           string                 `json:"image"`
	ImageURL        string                 `json:"imageUrl"`
	Locale          string                 `json:"locale"`
	Name            string                 `json:"name"`
	Phone           string                 `json:"phone"`
	PrimaryWalletID string                 `json:"primaryWalletId"`
	ReferralCode    string                 `json:"referralCode"`
	ReferralUserID  string                 `json:"referralUserId"`
	Roles           []string               `json:"roles"`
	UpdatedAt       time.Time              `json:"updatedAt"`
	Verified        bool                   `json:"verified"`
}

type UpdateUserDataBody struct {
	Name       string      `json:"name,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	CustomData interface{} `json:"customData,omitempty"`
}

type CreateUserDataBody struct {
	Email           string      `json:"email"`
	Code            string      `json:"code,omitempty"`
	Locale          string      `json:"locale,omitempty"`
	Name            string      `json:"name,omitempty"`
	Password        string      `json:"password,"`
	Approved        bool        `json:"approved"`
	Blocked         bool        `json:"blocked"`
	CustomData      interface{} `json:"customData,omitempty"`
	PrimaryWalletID string      `json:"primaryWalletId,omitempty"`
	Verified        bool        `json:"verified"`
}
