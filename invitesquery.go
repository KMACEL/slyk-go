package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetInvites is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites
func (c Client) GetInvites(filter ...*getPaymentMethodFilter) (*Invites, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkInvites)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var invites Invites
	errUnmarshal := json.Unmarshal(resp.Body(), &invites)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invites, nil
}

// GetInviteWithCode is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-code
func (c Client) GetInviteWithCode(inviteCode string) (*Invite, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkInvites + "/" + inviteCode)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(resp.Body(), &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}

// GetInviteWithCodeForValidate is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-code-validate
func (c Client) GetInviteWithCodeForValidate(inviteCode string) (*InviteForValidate, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkInvites + "/" + inviteCode + "/validate")

	fmt.Println(string(resp.Body()))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var inviteForValidate InviteForValidate
	errUnmarshal := json.Unmarshal(resp.Body(), &inviteForValidate)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &inviteForValidate, nil
}

// CreateInvite is
// https://developers.slyk.io/slyk/reference/endpoints#post-invites
func (c Client) CreateInvite(createInviteDataBody *CreateInviteDataBody) (*Invite, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(createInviteDataBody).
		Post(linkInvites)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(resp.Body(), &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}

// CancelInvite is
// https://developers.slyk.io/slyk/reference/endpoints#post-invites-code-cancel
func (c Client) CancelInvite(inviteCode string) (*Invite, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Post(linkInvites + "/" + inviteCode + "/cancel")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(resp.Body(), &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}

// SendInvite is
// https://developers.slyk.io/slyk/reference/endpoints#post-invites-send
// TODO : çalışmıyor
func (c Client) SendInvite(sendInviteDataBody *SendInviteDataBody) (*Invite, error) {
	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		SetBody(sendInviteDataBody).
		Post(linkInvites + "/send")

	fmt.Println(string(resp.Body()))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var invite Invite
	errUnmarshal := json.Unmarshal(resp.Body(), &invite)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &invite, nil
}
