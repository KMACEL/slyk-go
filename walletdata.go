package slyk

import (
	"strconv"
	"strings"
	"time"
)

type getWalletFilter map[string]string

type Wallets struct {
	Data  []WalletData `json:"data"`
	Total int          `json:"total"`
}

type WalletData struct {
	CreatedAt  time.Time `json:"createdAt"`
	CustomData struct {
	} `json:"customData"`
	Description interface{} `json:"description"`
	ID          string      `json:"id"`
	Locked      bool        `json:"locked"`
	Metadata    struct {
	} `json:"metadata"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"ownerId"`
	Reference string    `json:"reference"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetWalletFilter() *getWalletFilter {
	return &getWalletFilter{}
}

func (g *getWalletFilter) SetIDForIN(id ...string) *getWalletFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getWalletFilter) SetIDForNIN(id ...string) *getWalletFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getWalletFilter) SetLocked(locked bool) *getWalletFilter {
	(*g)["filter[locked]"] = strconv.FormatBool(locked)
	return g
}

func (g *getWalletFilter) SetName(name string) *getWalletFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getWalletFilter) SetOwnerID(ownerId string) *getWalletFilter {
	(*g)["filter[ownerId]"] = ownerId
	return g
}

func (g *getWalletFilter) SetReference(reference string) *getWalletFilter {
	(*g)["filter[reference]"] = reference
	return g
}

func (g *getWalletFilter) SetSortWithCreatedAt() *getWalletFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletFilter) SetSortWithCreatedAtReverse() *getWalletFilter {
	(*g)["sort"] = "-createdAt"
	return g
}
