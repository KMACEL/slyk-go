package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getMovementsFilter map[string]string
type getMovementWithIDFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ███╗   ███╗ ██████╗ ██╗   ██╗███████╗███╗   ███╗███████╗███╗   ██╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ████╗ ████║██╔═══██╗██║   ██║██╔════╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██╔████╔██║██║   ██║██║   ██║█████╗  ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetMovementFilter is
func GetMovementsFilter() *getMovementsFilter {
	return &getMovementsFilter{}
}

func (g *getMovementsFilter) SetGenericQueryParameter(key string, value interface{}) *getMovementsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getMovementsFilter) SetAssetCode(assetCode string) *getMovementsFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getMovementsFilter) SetAssetCodeWithIN(assetCode ...string) *getMovementsFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getMovementsFilter) SetAssetCodeWithNIN(assetCode ...string) *getMovementsFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

// Format : 2019-07-21
func (g *getMovementsFilter) SetCreatedAtWithGTE(date string) *getMovementsFilter {
	(*g)["filter[createdAt]"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getMovementsFilter) SetCreatedAtWithLTE(date string) *getMovementsFilter {
	(*g)["filter[createdAt]"] = "lte:" + date
	return g
}

func (g *getMovementsFilter) SetID(id string) *getMovementsFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getMovementsFilter) SetIDWithIN(id ...string) *getMovementsFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getMovementsFilter) SetIDWithNIN(id ...string) *getMovementsFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getMovementsFilter) SetTransaction(transactionID string) *getMovementsFilter {
	(*g)["filter[transactionId]"] = transactionID
	return g
}

func (g *getMovementsFilter) SetTransactionIDWithIN(transactionID ...string) *getMovementsFilter {
	(*g)["filter[transactionId]"] = "in:" + strings.Join(transactionID, ",")
	return g
}

func (g *getMovementsFilter) SetTransactionIDWithNIN(transactionID ...string) *getMovementsFilter {
	(*g)["filter[transactionId]"] = "nin:" + strings.Join(transactionID, ",")
	return g
}

func (g *getMovementsFilter) SetWalletID(walletID string) *getMovementsFilter {
	(*g)["filter[walletId]"] = walletID
	return g
}

func (g *getMovementsFilter) SetWalletIDWithIN(walletID ...string) *getMovementsFilter {
	(*g)["filter[walletId]"] = "in:" + strings.Join(walletID, ",")
	return g
}

func (g *getMovementsFilter) SetWalletIDWithNIN(walletID ...string) *getMovementsFilter {
	(*g)["filter[walletId]"] = "nin:" + strings.Join(walletID, ",")
	return g
}

func (g *getMovementsFilter) SetSortWithCreatedAt() *getMovementsFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getMovementsFilter) SetSortWithCreatedAtReverse() *getMovementsFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getMovementsFilter) SetSortWithAmount(amount string) *getMovementsFilter {
	(*g)["filter[amount]"] = "amount"
	return g
}

func (g *getMovementsFilter) SetSortWithAmountReverse(amount string) *getMovementsFilter {
	(*g)["filter[amount]"] = "-amount"
	return g
}

func (g *getMovementsFilter) SetAvailableTransactionWithUser() *getMovementsFilter {
	(*g)["include"] = "transaction"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getMovementsFilter) SetPageSize(size int) *getMovementsFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getMovementsFilter) SetPageNumber(number int) *getMovementsFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
 ██████╗ ███████╗████████╗        ███╗   ███╗ ██████╗ ██╗   ██╗███████╗███╗   ███╗███████╗███╗   ██╗████████╗        ██╗    ██╗██╗████████╗██╗  ██╗        ██╗██████╗         ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ████╗ ████║██╔═══██╗██║   ██║██╔════╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ██║    ██║██║╚══██╔══╝██║  ██║        ██║██╔══██╗        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██╔████╔██║██║   ██║██║   ██║█████╗  ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           ██║ █╗ ██║██║   ██║   ███████║        ██║██║  ██║        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██║███╗██║██║   ██║   ██╔══██║        ██║██║  ██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ╚███╔███╔╝██║   ██║   ██║  ██║        ██║██████╔╝        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝            ╚══╝╚══╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝╚═════╝         ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝

*/

// GetMovementWithIDFilter is
func GetMovementWithIDFilter() *getMovementWithIDFilter {
	return &getMovementWithIDFilter{}
}

func (g *getMovementWithIDFilter) SetGenericQueryParameter(key string, value interface{}) *getMovementWithIDFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getMovementWithIDFilter) SetAvailableTransactionWithUser() *getMovementWithIDFilter {
	(*g)["include"] = "transaction"
	return g
}
