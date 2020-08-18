package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getInvitesFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██╗███╗   ██╗██╗   ██╗██╗████████╗███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║████╗  ██║██║   ██║██║╚══██╔══╝██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║██╔██╗ ██║██║   ██║██║   ██║   █████╗  ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║██║╚██╗██║╚██╗ ██╔╝██║   ██║   ██╔══╝  ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║██║ ╚████║ ╚████╔╝ ██║   ██║   ███████╗███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝╚═╝  ╚═══╝  ╚═══╝  ╚═╝   ╚═╝   ╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetInvitesFilter() *getInvitesFilter {
	return &getInvitesFilter{}
}

func (g *getInvitesFilter) SetGenericQueryParameter(key string, value interface{}) *getInvitesFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getInvitesFilter) SetCode(code string) *getInvitesFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getInvitesFilter) SetCodeWithIN(code ...string) *getInvitesFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getInvitesFilter) SetCodeWithNIN(code ...string) *getInvitesFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

// Format : 2019-07-21
func (g *getInvitesFilter) SetExpiredAtWithGTE(date string) *getInvitesFilter {
	(*g)["filter[expiredAt]"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getInvitesFilter) SetExpiredAtWithLTE(date string) *getInvitesFilter {
	(*g)["filter[expiredAt]"] = "lte:" + date
	return g
}

func (g *getInvitesFilter) SetInvitedEmail(invitedEmail string) *getInvitesFilter {
	(*g)["filter[invitedEmail]"] = invitedEmail
	return g
}

func (g *getInvitesFilter) SetInvitedEmailWithIN(invitedEmail ...string) *getInvitesFilter {
	(*g)["filter[invitedEmail]"] = "in:" + strings.Join(invitedEmail, ",")
	return g
}

func (g *getInvitesFilter) SetInvitedEmailWithNIN(invitedEmail ...string) *getInvitesFilter {
	(*g)["filter[invitedEmail]"] = "nin:" + strings.Join(invitedEmail, ",")
	return g
}

func (g *getInvitesFilter) SetInvitedUserID(invitedUserId string) *getInvitesFilter {
	(*g)["filter[invitedUserId]"] = invitedUserId
	return g
}

func (g *getInvitesFilter) SetInvitedUserIDWithIN(invitedUserID ...string) *getInvitesFilter {
	(*g)["filter[invitedUserId]"] = "in:" + strings.Join(invitedUserID, ",")
	return g
}

func (g *getInvitesFilter) SetInvitedUserIDWithNIN(invitedUserID ...string) *getInvitesFilter {
	(*g)["filter[invitedUserId]"] = "nin:" + strings.Join(invitedUserID, ",")
	return g
}

func (g *getInvitesFilter) SetInviterUserID(inviterUserID string) *getInvitesFilter {
	(*g)["filter[inviterUserId]"] = inviterUserID
	return g
}

func (g *getInvitesFilter) SetInviterUserIDWithIN(inviterUserID ...string) *getInvitesFilter {
	(*g)["filter[inviterUserId]"] = "in:" + strings.Join(inviterUserID, ",")
	return g
}

func (g *getInvitesFilter) SetInviterUserIDWithNIN(inviterUserID ...string) *getInvitesFilter {
	(*g)["filter[inviterUserId]"] = "nin:" + strings.Join(inviterUserID, ",")
	return g
}

func (g *getInvitesFilter) SetStatus(status string) *getInvitesFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *getInvitesFilter) SetStatusWithIN(status ...string) *getInvitesFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *getInvitesFilter) SetStatusWithNIN(status ...string) *getInvitesFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *getInvitesFilter) SetType(typeInvite string) *getInvitesFilter {
	(*g)["filter[type]"] = typeInvite
	return g
}

func (g *getInvitesFilter) SetSortWithCreatedAt() *getInvitesFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getInvitesFilter) SetSortWithCreatedAtReverse() *getInvitesFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getInvitesFilter) SetSortWithExpiredAt() *getInvitesFilter {
	(*g)["sort"] = "expiredAt"
	return g
}

func (g *getInvitesFilter) SetSortWithExpiredAtReverse() *getInvitesFilter {
	(*g)["sort"] = "-expiredAt"
	return g
}

func (g *getInvitesFilter) SetSortWithUpdatedAt() *getInvitesFilter {
	(*g)["sort"] = "updatedAt"
	return g
}

func (g *getInvitesFilter) SetSortWithUpdatedAtReverse() *getInvitesFilter {
	(*g)["sort"] = "-updatedAt"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getInvitesFilter) SetPageSize(size int) *getInvitesFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getInvitesFilter) SetPageNumber(number int) *getInvitesFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
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
func CreateInviteDataForBody() *CreateInviteDataBody {
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
