package slyk

import (
	"strconv"
	"strings"
)

type getUserFilter map[string]string

type PatchUser struct {
	Name       string      `json:"name"`
	Locale     string      `json:"locale"`
	CustomData interface{} `json:"customData"`
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
