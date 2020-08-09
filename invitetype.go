package slyk

import "time"

type Invites struct {
	Data  []InviteData `json:"data"`
	Total int          `json:"total"`
}

type Invite struct {
	Data InviteData `json:"data"`
}

type InviteData struct {
	Code          string      `json:"code"`
	CreatedAt     time.Time   `json:"createdAt"`
	ExpiredAt     time.Time   `json:"expiredAt"`
	InvitedEmail  string      `json:"invitedEmail"`
	InvitedUserID interface{} `json:"invitedUserId"`
	InviterUserID interface{} `json:"inviterUserId"`
	Metadata      interface{} `json:"metadata"`
	Status        string      `json:"status"`
	Type          string      `json:"type"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	URL           string      `json:"url"`
}

type InviteForValidate struct {
	Data struct {
		ReferrerName string `json:"referrerName"`
	} `json:"data"`
}

type CreateInviteDataBody struct {
	Email         string `json:"email,omitempty"`
	InviterUserID string `json:"inviterUserId,omitempty"`
}

type SendInviteDataBody struct {
	Email         []string `json:"email"`
	InviterUserID string   `json:"inviterUserId,omitempty"`
}
