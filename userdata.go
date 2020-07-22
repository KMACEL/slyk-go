package slyk

import (
	"strconv"
	"strings"
	"time"
)

type Users struct {
	Data  []Data `json:"data"`
	Total int    `json:"total"`
}

type User struct {
	Data `json:"data"`
}

type Data struct {
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

type getUserFilter map[string]string

type PatchUser struct {
	Name       string      `json:"name"`
	Locale     string      `json:"locale"`
	CustomData interface{} `json:"customData"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func GetUserFilter() *getUserFilter {
	return &getUserFilter{}
}

func (u *getUserFilter) SetApproved(approved bool) *getUserFilter {
	(*u)["filter[approved]"] = strconv.FormatBool(approved)
	return u
}

func (u *getUserFilter) SetBlocked(blocked bool) *getUserFilter {
	(*u)["filter[blocked]"] = strconv.FormatBool(blocked)
	return u
}

func (u *getUserFilter) SetEmail(email string) *getUserFilter {
	(*u)["filter[email]"] = email
	return u
}

func (u *getUserFilter) SetID(id ...string) *getUserFilter {
	// TODO YANLIŞ OLAN VARSA HİÇ GELMİYOR
	(*u)["filter[id]"] = "in:" + strings.Join(id, ",")
	return u
}

func (u *getUserFilter) SetName(name string) *getUserFilter {
	(*u)["filter[name]"] = name
	return u
}

func (u *getUserFilter) SetPhone(phone string) *getUserFilter {
	(*u)["filter[phone]"] = phone
	return u
}

func (u *getUserFilter) SetPrimaryWalletId(primaryWalletId string) *getUserFilter {
	(*u)["filter[primaryWalletId]"] = primaryWalletId
	return u
}

func (u *getUserFilter) SetReferralCode(referralCode string) *getUserFilter {
	(*u)["filter[referralCode]"] = referralCode
	return u
}

func (u *getUserFilter) SetReferralUserID(referralUserId string) *getUserFilter {
	(*u)["filter[referralUserId]"] = referralUserId
	return u
}

func (u *getUserFilter) SetRole(role string) *getUserFilter {
	(*u)["filter[role]"] = role
	return u
}

func (u *getUserFilter) SetVerified(verified bool) *getUserFilter {
	(*u)["filter[verified]"] = strconv.FormatBool(verified)
	return u
}

func (u *getUserFilter) SetSortWithCreatedAt() *getUserFilter {
	(*u)["sort"] = "createdAt"
	return u
}

func (u *getUserFilter) SetSortWithCreatedAtReverse() *getUserFilter {
	(*u)["sort"] = "-createdAt"
	return u
}

func (u *getUserFilter) SetSortWithUpdatedAt() *getUserFilter {
	(*u)["sort"] = "updatedAt"
	return u
}

func (u *getUserFilter) SetSortWithUpdatedAtReverse() *getUserFilter {
	(*u)["sort"] = "-updatedAt"
	return u
}

// Defines the number of results per page. Default = 30.
func (u *getUserFilter) SetPageSize(size int) *getUserFilter {
	(*u)["page[size]"] = strconv.Itoa(size)
	return u
}

// Defines the number of the page to retrieve. Default = 1
func (u *getUserFilter) SetPageNumber(number int) *getUserFilter {
	(*u)["page[number]"] = strconv.Itoa(number)
	return u
}

func CreatePatchData() *PatchUser {
	return &PatchUser{}
}

func (p *PatchUser) SetName(name string) *PatchUser {
	p.Name = name
	return p
}

func (p *PatchUser) SetLocale(locale string) *PatchUser {
	p.Locale = locale
	return p
}

// TODO Çalışmıyor bakılacak
func (p *PatchUser) SetCustomData(customData interface{}) *PatchUser {
	p.CustomData = customData
	return p
}

func CreatePostUser() *CreateUser {
	return &CreateUser{}
}

func (c *CreateUser) SetName(name string) *CreateUser {
	c.Name = name
	return c
}

func (c *CreateUser) SetEmail(email string) *CreateUser {
	c.Email = email
	return c
}

func (c *CreateUser) SetPassword(password string) *CreateUser {
	c.Password = password
	return c
}

func (c *CreateUser) SetLocale(locale string) *CreateUser {
	c.Locale = locale
	return c
}
