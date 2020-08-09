package slyk

import (
	"encoding/json"
)

// GetUser is returns the Slyk user list.
// https://developers.slyk.io/slyk/reference/endpoints#get-users
func (c Client) GetUsers(filter ...*getUserFilter) (*Users, error) {
	getBody, err := c.genericGetQuery(linkUsers, merge(filter))
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

// GetUserWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-users-id
func (c Client) GetUserWithID(userID string) (*User, error) {
	getBody, err := c.genericGetQuery(linkUsers+"/"+userID, nil)
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
func (c Client) UpdateUser(userID string, updateUserData *UpdateUserData) (*User, error) {
	getBody, err := c.genericPatchQuery(linkUsers+"/"+userID, updateUserData)
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
// TODO Diğer parametrelere bakılacak
// TODO Password tipi kontrol edilecek
func (c Client) CreateUser(createUserdata *CreateUserData) (*User, error) {
	getBody, err := c.genericPostQuery(linkUsers, createUserdata)
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
	_, err := c.genericPostQuery(linkUsers+"/"+userID+approve, nil)
	return err
}

// SetUserBlock is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-block
func (c Client) SetUserBlock(userID string) error {
	_, err := c.genericPostQuery(linkUsers+"/"+userID+block, nil)
	return err
}

// SetUserUnblock is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-unblock
func (c Client) SetUserUnblock(userID string) error {
	_, err := c.genericPostQuery(linkUsers+"/"+userID+unblock, nil)
	return err
}

// ChangePassword is
// https://developers.slyk.io/slyk/reference/endpoints#post-users-id-change-password
func (c Client) ChangePassword(userID string, psw string) error {
	_, err := c.genericPostQuery(linkUsers+"/"+userID+changePassword, struct {
		Password string `json:"password"`
	}{
		Password: psw,
	})
	return err
}
