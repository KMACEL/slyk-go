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
	Approved   bool      `json:"approved"`
	Blocked    bool      `json:"blocked"`
	CreatedAt  time.Time `json:"createdAt"`
	CustomData struct {
	} `json:"customData"`
	Email           string      `json:"email"`
	ID              string      `json:"id"`
	Image           interface{} `json:"image"`
	ImageURL        interface{} `json:"imageUrl"`
	Locale          string      `json:"locale"`
	Name            string      `json:"name"`
	Phone           interface{} `json:"phone"`
	PrimaryWalletID string      `json:"primaryWalletId"`
	ReferralCode    string      `json:"referralCode"`
	ReferralUserID  interface{} `json:"referralUserId"`
	Roles           []string    `json:"roles"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Verified        bool        `json:"verified"`
}

type UpdateUserData struct {
	Name       string      `json:"name"`
	Locale     string      `json:"locale"`
	CustomData interface{} `json:"customData"`
}

type CreateUserData struct {
	Email    string `json:"email"`
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
