package slyk

import (
	"encoding/json"
)

// GetInvites is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites
func (c Client) GetInvites(filter ...*getInviteFilter) (*Invites, error) {
	getBody, err := c.genericGetQuery(linkInvites, merge(filter))
	if err != nil {
		return nil, err
	}

	var invites Invites
	errUnmarshal := json.Unmarshal(getBody, &invites)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invites, nil
}

// GetInviteWithCode is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-code
func (c Client) GetInviteWithCode(inviteCode string) (*Invite, error) {
	getBody, err := c.genericGetQuery(linkInvites+"/"+inviteCode, nil)
	if err != nil {
		return nil, err
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(getBody, &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}

// GetInviteWithCodeForValidate is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-code-validate
func (c Client) GetInviteWithCodeForValidate(inviteCode string) (*InviteForValidate, error) {
	getBody, err := c.genericGetQuery(linkInvites+"/"+inviteCode+"/validate", nil)
	if err != nil {
		return nil, err
	}

	var inviteForValidate InviteForValidate
	errUnmarshal := json.Unmarshal(getBody, &inviteForValidate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &inviteForValidate, nil
}

// CreateInvite is
// https://developers.slyk.io/slyk/reference/endpoints#post-invites
func (c Client) CreateInvite(createInviteDataBody *CreateInviteDataBody) (*Invite, error) {
	getBody, err := c.genericPostQuery(linkInvites, createInviteDataBody)
	if err != nil {
		return nil, err
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(getBody, &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}

// CancelInvite is
// https://developers.slyk.io/slyk/reference/endpoints#post-invites-code-cancel
func (c Client) CancelInvite(inviteCode string) (*Invite, error) {
	getBody, err := c.genericPostQuery(linkInvites+"/"+inviteCode+"/cancel", nil)
	if err != nil {
		return nil, err
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(getBody, &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}

// SendInvite is
// https://developers.slyk.io/slyk/reference/endpoints#post-invites-send
// TODO : çalışmıyor
func (c Client) SendInvite(sendInviteDataBody *SendInviteDataBody) (*Invite, error) {
	getBody, err := c.genericPostQuery(linkInvites+"/send", sendInviteDataBody)
	if err != nil {
		return nil, err
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(getBody, &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}
