package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getInviteFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██╗███╗   ██╗██╗   ██╗██╗████████╗███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║████╗  ██║██║   ██║██║╚══██╔══╝██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║██╔██╗ ██║██║   ██║██║   ██║   █████╗  ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║██║╚██╗██║╚██╗ ██╔╝██║   ██║   ██╔══╝  ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║██║ ╚████║ ╚████╔╝ ██║   ██║   ███████╗███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝╚═╝  ╚═══╝  ╚═══╝  ╚═╝   ╚═╝   ╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetInvitesFilter() *getInviteFilter {
	return &getInviteFilter{}
}

func (g *getInviteFilter) SetGenericQueryParameter(key string, value interface{}) *getInviteFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getInviteFilter) SetCode(code string) *getInviteFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getInviteFilter) SetCodeWithIN(code ...string) *getInviteFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getInviteFilter) SetCodeWithNIN(code ...string) *getInviteFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

// Format : 2019-07-21
func (g *getMovementFilter) SetExpiredAtWithGTE(date string) *getMovementFilter {
	(*g)["expiredAt"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getMovementFilter) SetExpiredAtWithLTE(date string) *getMovementFilter {
	(*g)["expiredAt"] = "lte:" + date
	return g
}

func (g *getInviteFilter) SetInvitedEmail(invitedEmail string) *getInviteFilter {
	(*g)["filter[invitedEmail]"] = invitedEmail
	return g
}

func (g *getInviteFilter) SetInvitedEmailWithIN(invitedEmail ...string) *getInviteFilter {
	(*g)["filter[invitedEmail]"] = "in:" + strings.Join(invitedEmail, ",")
	return g
}

func (g *getInviteFilter) SetInvitedEmailWithNIN(invitedEmail ...string) *getInviteFilter {
	(*g)["filter[invitedEmail]"] = "nin:" + strings.Join(invitedEmail, ",")
	return g
}

func (g *getInviteFilter) SetInvitedUserID(invitedEmail string) *getInviteFilter {
	(*g)["filter[invitedUserId]"] = invitedEmail
	return g
}

func (g *getInviteFilter) SetInvitedUserIDWithIN(invitedUserID ...string) *getInviteFilter {
	(*g)["filter[invitedUserId]"] = "in:" + strings.Join(invitedUserID, ",")
	return g
}

func (g *getInviteFilter) SetInvitedUserIDWithNIN(invitedUserID ...string) *getInviteFilter {
	(*g)["filter[invitedUserId]"] = "nin:" + strings.Join(invitedUserID, ",")
	return g
}

func (g *getInviteFilter) SetInviterUserID(inviterUserID string) *getInviteFilter {
	(*g)["filter[inviterUserId]"] = inviterUserID
	return g
}

func (g *getInviteFilter) SetInviterUserIDWithIN(inviterUserID ...string) *getInviteFilter {
	(*g)["filter[inviterUserId]"] = "in:" + strings.Join(inviterUserID, ",")
	return g
}

func (g *getInviteFilter) SetInviterUserIDWithNIN(inviterUserID ...string) *getInviteFilter {
	(*g)["filter[inviterUserId]"] = "nin:" + strings.Join(inviterUserID, ",")
	return g
}

func (g *getInviteFilter) SetStatus(status string) *getInviteFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *getInviteFilter) SetStatusWithIN(status ...string) *getInviteFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *getInviteFilter) SetStatusWithNIN(status ...string) *getInviteFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *getInviteFilter) SetType(typeInvite string) *getInviteFilter {
	(*g)["filter[type]"] = typeInvite
	return g
}

func (g *getInviteFilter) SetSortWithCreatedAt() *getInviteFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getInviteFilter) SetSortWithCreatedAtReverse() *getInviteFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getInviteFilter) SetSortWithExpiredAt() *getInviteFilter {
	(*g)["sort"] = "expiredAt"
	return g
}

func (g *getInviteFilter) SetSortWithExpiredAtReverse() *getInviteFilter {
	(*g)["sort"] = "-expiredAt"
	return g
}

func (g *getInviteFilter) SetSortWithUpdatedAt() *getInviteFilter {
	(*g)["sort"] = "updatedAt"
	return g
}

func (g *getInviteFilter) SetSortWithUpdatedAtReverse() *getInviteFilter {
	(*g)["sort"] = "-updatedAt"
	return g
}

// Defines the number of results per page. Default = 30.
func (u *getInviteFilter) SetPageSize(size int) *getInviteFilter {
	(*u)["page[size]"] = strconv.Itoa(size)
	return u
}

// Defines the number of the page to retrieve. Default = 1
func (u *getInviteFilter) SetPageNumber(number int) *getInviteFilter {
	(*u)["page[number]"] = strconv.Itoa(number)
	return u
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ██╗███╗   ██╗██╗   ██╗██╗████████╗███████╗███████╗        ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██║████╗  ██║██║   ██║██║╚══██╔══╝██╔════╝██╔════╝        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║██╔██╗ ██║██║   ██║██║   ██║   █████╗  ███████╗        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║██║╚██╗██║╚██╗ ██╔╝██║   ██║   ██╔══╝  ╚════██║        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ██║██║ ╚████║ ╚████╔╝ ██║   ██║   ███████╗███████║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝╚═╝  ╚═══╝  ╚═══╝  ╚═╝   ╚═╝   ╚══════╝╚══════╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CreateInvitesBody is
func CreateInvitesBody() *CreateInviteDataBody {
	return &CreateInviteDataBody{}
}

func (c *CreateInviteDataBody) SetEmail(email string) *CreateInviteDataBody {
	c.Email = email
	return c
}

func (c *CreateInviteDataBody) SetInviterUserID(inviterUserID string) *CreateInviteDataBody {
	c.InviterUserID = inviterUserID
	return c
}
