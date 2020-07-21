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
