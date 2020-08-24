package slyk

import (
	"encoding/json"
)

// GetUser is returns the Slyk user list.
// https://developers.slyk.io/slyk/reference/endpoints#get-users
func (c Client) GetUsers(filter ...*getUsersFilter) (*Users, error) {
	getBody, err := c.GenericGetQuery(linkUsers, merge(filter))
	if err != nil {
		return nil, err
	}

	var users Users
	errUnmarshal := json.Unmarshal(getBody, &users)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &users, nil
}

// GetUserWithID is Fetches information about a user. User Id takes as parameter.
// https://developers.slyk.io/slyk/reference/endpoints#get-users-id
func (c Client) GetUserWithID(userID string) (*User, error) {
	getBody, err := c.GenericGetQuery(linkUsers+"/"+userID, nil)
	if err != nil {
		return nil, err
	}

	var user User
	errUnmarshal := json.Unmarshal(getBody, &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}

// UpdateUser is
// https://developers.slyk.io/slyk/reference/endpoints#patch-users-id
func (c Client) UpdateUser(userID string, updateUserData *UpdateUserDataBody) (*User, error) {
	getBody, err := c.GenericPatchQuery(linkUsers+"/"+userID, updateUserData)
	if err != nil {
		return nil, err
	}

	var user User
	errUnmarshal := json.Unmarshal(getBody, &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}

// CreateUser is
// https://developers.slyk.io/slyk/reference/endpoints#post-users
func (c Client) CreateUser(createUserdata *CreateUserDataBody) (*User, error) {
	getBody, err := c.GenericPostQuery(linkUsers, createUserdata)
	if err != nil {
		return nil, err
	}

	var user User
	errUnmarshal := json.Unmarshal(getBody, &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &user, nil
}

// SetUserApprove is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-approve
func (c Client) SetUserApprove(userID string) error {
	_, err := c.GenericPostQuery(linkUsers+"/"+userID+approve, nil)
	return err
}

// SetUserBlock is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-block
func (c Client) SetUserBlock(userID string) error {
	_, err := c.GenericPostQuery(linkUsers+"/"+userID+block, nil)
	return err
}

// SetUserUnblock is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-unblock
func (c Client) SetUserUnblock(userID string) error {
	_, err := c.GenericPostQuery(linkUsers+"/"+userID+unblock, nil)
	return err
}

// ChangePassword is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-change-password
func (c Client) ChangePassword(userID string, psw string) error {
	_, err := c.GenericPostQuery(linkUsers+"/"+userID+changePassword, struct {
		Password string `json:"password"`
	}{
		Password: psw,
	})
	return err
}
