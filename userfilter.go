package slyk

import (
	"strconv"
	"strings"
)

type userFilter map[string]string

func CreateUserFilter() *userFilter {
	return &userFilter{}
}

func (u *userFilter) SetApproved(approved bool) *userFilter {
	(*u)["filter[approved]"] = strconv.FormatBool(approved)
	return u
}

func (u *userFilter) SetBlocked(blocked bool) *userFilter {
	(*u)["filter[blocked]"] = strconv.FormatBool(blocked)
	return u
}

func (u *userFilter) SetEmail(email string) *userFilter {
	(*u)["filter[email]"] = email
	return u
}

func (u *userFilter) SetID(id ...string) *userFilter {
	// TODO YANLIŞ OLAN VARSA HİÇ GELMİYOR
	(*u)["filter[id]"] = "in:" + strings.Join(id, ",")
	return u
}

func (u *userFilter) SetName(name string) *userFilter {
	(*u)["filter[name]"] = name
	return u
}

func (u *userFilter) SetPhone(phone string) *userFilter {
	(*u)["filter[phone]"] = phone
	return u
}

func (u *userFilter) SetPrimaryWalletId(primaryWalletId string) *userFilter {
	(*u)["filter[primaryWalletId]"] = primaryWalletId
	return u
}

func (u *userFilter) SetReferralCode(referralCode string) *userFilter {
	(*u)["filter[referralCode]"] = referralCode
	return u
}

func (u *userFilter) SetReferralUserID(referralUserId string) *userFilter {
	(*u)["filter[referralUserId]"] = referralUserId
	return u
}

func (u *userFilter) SetRole(role string) *userFilter {
	(*u)["filter[role]"] = role
	return u
}

func (u *userFilter) SetVerified(verified bool) *userFilter {
	(*u)["filter[verified]"] = strconv.FormatBool(verified)
	return u
}

func (u *userFilter) SetSortWithCreatedAt() *userFilter {
	(*u)["sort"] = "createdAt"
	return u
}

func (u *userFilter) SetSortWithCreatedAtReverse() *userFilter {
	(*u)["sort"] = "-createdAt"
	return u
}

func (u *userFilter) SetSortWithUpdatedAt() *userFilter {
	(*u)["sort"] = "updatedAt"
	return u
}

func (u *userFilter) SetSortWithUpdatedAtReverse() *userFilter {
	(*u)["sort"] = "-updatedAt"
	return u
}

// Defines the number of results per page. Default = 30.
func (u *userFilter) SetPageSize(size int) *userFilter {
	(*u)["page[size]"] = strconv.Itoa(size)
	return u
}

// Defines the number of the page to retrieve. Default = 1
func (u *userFilter) SetPageNumber(number int) *userFilter {
	(*u)["page[number]"] = strconv.Itoa(number)
	return u
}
