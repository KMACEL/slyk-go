package slyk

import (
	"strconv"
	"strings"
)

type getMovementFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ███╗   ███╗ ██████╗ ██╗   ██╗███████╗███╗   ███╗███████╗███╗   ██╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ████╗ ████║██╔═══██╗██║   ██║██╔════╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██╔████╔██║██║   ██║██║   ██║█████╗  ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetMovementFilter is
func GetMovementFilter() *getMovementFilter {
	return &getMovementFilter{}
}

// Format : 2019-07-21
func (g *getMovementFilter) SetCreatedAtWithGTE(date string) *getMovementFilter {
	(*g)["createdAt"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getMovementFilter) SetCreatedAtWithLTE(date string) *getMovementFilter {
	(*g)["createdAt"] = "lte:" + date
	return g
}

func (g *getMovementFilter) SetWalletIDWithIN(walletID ...string) *getMovementFilter {
	(*g)["filter[walletId]"] = "in:" + strings.Join(walletID, ",")
	return g
}

func (g *getMovementFilter) SetWalletIDWithNIN(walletID ...string) *getMovementFilter {
	(*g)["filter[walletId]"] = "nin:" + strings.Join(walletID, ",")
	return g
}

// asset btc,eth,ltc
func (g *getMovementFilter) SetAssetCodeWithIN(assetCode string) *getMovementFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getMovementFilter) SetAssetCodeWithNIN(assetCode string) *getMovementFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
	return g
}

func (g *getMovementFilter) SetIDWithIN(id ...string) *getMovementFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getMovementFilter) SetIDWithNIN(id ...string) *getMovementFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getMovementFilter) SetTransactionIDWithIN(transactionID string) *getMovementFilter {
	(*g)["filter[transactionId]"] = "in:" + transactionID
	return g
}

func (g *getMovementFilter) SetTransactionIDWithNIN(transactionID string) *getMovementFilter {
	(*g)["filter[transactionId]"] = "nin:" + transactionID
	return g
}

func (g *getMovementFilter) SetSortWithCreatedAt() *getMovementFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getMovementFilter) SetSortWithCreatedAtReverse() *getMovementFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getMovementFilter) SetSortWithAmount(amount string) *getMovementFilter {
	(*g)["filter[amount]"] = "amount"
	return g
}

func (g *getMovementFilter) SetSortWithAmountReverse(amount string) *getMovementFilter {
	(*g)["filter[amount]"] = "-amount"
	return g
}

func (g *getMovementFilter) SetAvailableTransactionWithUser() *getMovementFilter {
	(*g)["include"] = "transaction"
	return g
}

// Defines the number of results per page. Default = 30.
func (u *getMovementFilter) SetPageSize(size int) *getMovementFilter {
	(*u)["page[size]"] = strconv.Itoa(size)
	return u
}

// Defines the number of the page to retrieve. Default = 1
func (u *getMovementFilter) SetPageNumber(number int) *getMovementFilter {
	(*u)["page[number]"] = strconv.Itoa(number)
	return u
}
