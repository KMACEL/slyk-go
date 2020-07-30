package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetUser is
// https://developers.slyk.io/slyk/reference/endpoints#get-users
func (c Client) GetUser(filter ...*getUserFilter) (*Users, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkUsers)

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
// https://developers.slyk.io/slyk/reference/endpoints#get-users-id
func (c Client) GetUserWithID(userID string) (*User, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkUsers + "/" + userID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var user User
	errUnmarshal := json.Unmarshal(resp.Body(), &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}

// UpdateUser is
// https://developers.slyk.io/slyk/reference/endpoints#patch-users-id
func (c Client) UpdateUser(userID string, updateUserData *UpdateUserData) (*User, error) {
	resp, err := resty.New().R().
		SetBody(*updateUserData).
		SetHeader(headerApiKey, c.apiKey).
		Patch(linkUsers + "/" + userID)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var user User
	errUnmarshal := json.Unmarshal(resp.Body(), &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}

// CreateUser is
// https://developers.slyk.io/slyk/reference/endpoints#post-users
// TODO Diğer parametrelere bakılacak
// TODO Password tipi kontrol edilecek
func (c Client) CreateUser(createUserdata *CreateUserData) (*User, error) {
	resp, err := resty.New().R().
		SetBody(*createUserdata).
		SetHeader(headerApiKey, c.apiKey).
		Post(linkUsers)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var user User
	errUnmarshal := json.Unmarshal(resp.Body(), &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}

// SetUserApprove is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-approve
func (c Client) SetUserApprove(userID string) error {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Post(linkUsers + "/" + userID + approve)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	return nil
}

// SetUserBlock is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-block
func (c Client) SetUserBlock(userID string) error {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Post(linkUsers + "/" + userID + block)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	return nil
}

// SetUserUnblock is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-unblock
func (c Client) SetUserUnblock(userID string) error {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Post(linkUsers + "/" + userID + unblock)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	return nil
}

// ChangePassword is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-change-password
func (c Client) ChangePassword(userID string, psw string) error {
	resp, err := resty.New().R().
		SetBody(struct {
			Password string `json:"password"`
		}{
			Password: psw,
		}).
		SetHeader(headerApiKey, c.apiKey).
		Post(linkUsers + "/" + userID + changePassword)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	return nil
}
