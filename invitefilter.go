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

func (g *getInviteFilter) SetCodeWithIN(code string) *getInviteFilter {
	(*g)["filter[code]"] = "in:" + code
	return g
}

func (g *getInviteFilter) SetCodeWithNIN(code string) *getInviteFilter {
	(*g)["filter[code]"] = "nin:" + code
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

func (g *getInviteFilter) SetInvitedEmailWithIN(invitedEmail string) *getInviteFilter {
	(*g)["filter[invitedEmail]"] = "in:" + invitedEmail
	return g
}

func (g *getInviteFilter) SetInvitedEmailWithNIN(invitedEmail string) *getInviteFilter {
	(*g)["filter[invitedEmail]"] = "nin:" + invitedEmail
	return g
}

func (g *getInviteFilter) SetınvitedUserID(invitedUserId ...string) *getInviteFilter {
	(*g)["filter[invitedUserId]"] = strings.Join(invitedUserId, ",")
	return g
}

func (g *getInviteFilter) SetInviterUserIDWithIN(inviterUserId string) *getInviteFilter {
	(*g)["filter[inviterUserId]"] = "in:" + inviterUserId
	return g
}

func (g *getInviteFilter) SetInviterUserIDWithNIN(inviterUserId string) *getInviteFilter {
	(*g)["filter[inviterUserId]"] = "nin:" + inviterUserId
	return g
}

func (g *getInviteFilter) SetStatusWithIN(status string) *getInviteFilter {
	(*g)["filter[status]"] = "in:" + status
	return g
}

func (g *getInviteFilter) SetStatusWithNIN(status string) *getInviteFilter {
	(*g)["filter[status]"] = "nin:" + status
	return g
}

func (g *getInviteFilter) SetTypeWithNIN(typeInvite string) *getInviteFilter {
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
