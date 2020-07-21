package slyk

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Users struct {
	Data  []UserData `json:"data"`
	Total int        `json:"total"`
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

// GetUser is
func (c Client) GetUser(filter ...*userFilter) (*Users, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(*filter[0])
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkGetUsers)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}
	var users Users
	errUnmarshal := json.Unmarshal(resp.Body(), &users)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &users, nil
}

// GetUserWithID is
func (c Client) GetUserWithID(userID string) (*UserData, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkGetUsers + "/" + userID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var user UserData
	errUnmarshal := json.Unmarshal(resp.Body(), &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}
