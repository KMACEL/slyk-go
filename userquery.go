package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetUser is
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

// GetUserWithID is
func (c Client) PatchUser(userID string, patchUser *PatchUser) (*User, error) {
	resp, err := resty.New().R().
		SetBody(*patchUser).
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

// GetUserWithID is
// TODO Diğer parametrelere bakılacak
func (c Client) PostCreateUser(createUser *CreateUser) (*User, error) {
	resp, err := resty.New().R().
		SetBody(*createUser).
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
