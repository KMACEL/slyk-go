package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getUsersFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██╗   ██╗███████╗███████╗██████╗         ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║   ██║██╔════╝██╔════╝██╔══██╗        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║   ██║███████╗█████╗  ██████╔╝        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║   ██║╚════██║██╔══╝  ██╔══██╗        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚██████╔╝███████║███████╗██║  ██║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetUsersFilter() *getUsersFilter {
	return &getUsersFilter{}
}

func (g *getUsersFilter) SetGenericQueryParameter(key string, value interface{}) *getUsersFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getUsersFilter) SetApproved(approved bool) *getUsersFilter {
	(*g)["filter[approved]"] = strconv.FormatBool(approved)
	return g
}

func (g *getUsersFilter) SetBlocked(blocked bool) *getUsersFilter {
	(*g)["filter[blocked]"] = strconv.FormatBool(blocked)
	return g
}

func (g *getUsersFilter) SetEmail(email string) *getUsersFilter {
	(*g)["filter[email]"] = email
	return g
}

func (g *getUsersFilter) SetID(id string) *getUsersFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getUsersFilter) SetIDWithIN(id ...string) *getUsersFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getUsersFilter) SetName(name string) *getUsersFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getUsersFilter) SetPhone(phone string) *getUsersFilter {
	(*g)["filter[phone]"] = phone
	return g
}

func (g *getUsersFilter) SetPrimaryWalletId(primaryWalletId string) *getUsersFilter {
	(*g)["filter[primaryWalletId]"] = primaryWalletId
	return g
}

func (g *getUsersFilter) SetReferralCode(referralCode string) *getUsersFilter {
	(*g)["filter[referralCode]"] = referralCode
	return g
}

func (g *getUsersFilter) SetReferralUserID(referralUserID string) *getUsersFilter {
	(*g)["filter[referralUserId]"] = referralUserID
	return g
}

func (g *getUsersFilter) SetReferralUserIDWithIN(referralUserID ...string) *getUsersFilter {
	(*g)["filter[referralUserId]"] = "in:" + strings.Join(referralUserID, ",")
	return g
}

func (g *getUsersFilter) SetRole(role string) *getUsersFilter {
	(*g)["filter[role]"] = role
	return g
}

func (g *getUsersFilter) SetVerified(verified bool) *getUsersFilter {
	(*g)["filter[verified]"] = strconv.FormatBool(verified)
	return g
}

func (g *getUsersFilter) SetSortWithCreatedAt() *getUsersFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getUsersFilter) SetSortWithCreatedAtReverse() *getUsersFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getUsersFilter) SetSortWithUpdatedAt() *getUsersFilter {
	(*g)["sort"] = "updatedAt"
	return g
}

func (g *getUsersFilter) SetSortWithUpdatedAtReverse() *getUsersFilter {
	(*g)["sort"] = "-updatedAt"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getUsersFilter) SetPageSize(size int) *getUsersFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getUsersFilter) SetPageNumber(number int) *getUsersFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗        ██╗   ██╗███████╗███████╗██████╗         ██████╗  █████╗ ██████╗  █████╗ ███╗   ███╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██║   ██║██╔════╝██╔════╝██╔══██╗        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗████╗ ████║
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗          ██║   ██║███████╗█████╗  ██████╔╝        ██████╔╝███████║██████╔╝███████║██╔████╔██║
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝          ██║   ██║╚════██║██╔══╝  ██╔══██╗        ██╔═══╝ ██╔══██║██╔══██╗██╔══██║██║╚██╔╝██║
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗        ╚██████╔╝███████║███████╗██║  ██║        ██║     ██║  ██║██║  ██║██║  ██║██║ ╚═╝ ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝        ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝
*/

func UpdateUserDataForBody() *UpdateUserDataBody {
	return &UpdateUserDataBody{}
}

func (p *UpdateUserDataBody) SetName(name string) *UpdateUserDataBody {
	p.Name = name
	return p
}

func (p *UpdateUserDataBody) SetLocale(locale string) *UpdateUserDataBody {
	p.Locale = locale
	return p
}

func (p *UpdateUserDataBody) SetCustomData(customData interface{}) *UpdateUserDataBody {
	p.CustomData = customData
	return p
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ██╗   ██╗███████╗███████╗██████╗         ██████╗  █████╗ ██████╗  █████╗ ███╗   ███╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██║   ██║██╔════╝██╔════╝██╔══██╗        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗████╗ ████║
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║   ██║███████╗█████╗  ██████╔╝        ██████╔╝███████║██████╔╝███████║██╔████╔██║
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║   ██║╚════██║██╔══╝  ██╔══██╗        ██╔═══╝ ██╔══██║██╔══██╗██╔══██║██║╚██╔╝██║
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚██████╔╝███████║███████╗██║  ██║        ██║     ██║  ██║██║  ██║██║  ██║██║ ╚═╝ ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝        ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝
*/

// CreateUserParam is
func CreateUserDataForBody() *CreateUserDataBody {
	return &CreateUserDataBody{}
}

func (c *CreateUserDataBody) SetName(name string) *CreateUserDataBody {
	c.Name = name
	return c
}

func (c *CreateUserDataBody) SetEmail(email string) *CreateUserDataBody {
	c.Email = email
	return c
}

func (c *CreateUserDataBody) SetPassword(password string) *CreateUserDataBody {
	c.Password = password
	return c
}

func (c *CreateUserDataBody) SetLocale(locale string) *CreateUserDataBody {
	c.Locale = locale
	return c
}

func (c *CreateUserDataBody) SetCustomData(customData interface{}) *CreateUserDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateUserDataBody) SetApproved(approved bool) *CreateUserDataBody {
	c.Approved = approved
	return c
}

func (c *CreateUserDataBody) SetBlocked(blocked bool) *CreateUserDataBody {
	c.Blocked = blocked
	return c
}

func (c *CreateUserDataBody) SetCode(code string) *CreateUserDataBody {
	c.Code = code
	return c
}

func (c *CreateUserDataBody) SetPrimaryWalletID(primaryWalletID string) *CreateUserDataBody {
	c.PrimaryWalletID = primaryWalletID
	return c
}

func (c *CreateUserDataBody) SetVerified(verified bool) *CreateUserDataBody {
	c.Verified = verified
	return c
}
