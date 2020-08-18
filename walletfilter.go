package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getWalletsFilter map[string]string
type getWalletActivityWithIDFilter map[string]string
type getWalletBalanceWithIDFilter map[string]string
type getWalletMovementsFilter map[string]string
type getWalletTransactionstFilter map[string]string
type getWalletActivityFilter map[string]string
type getWalletBalanceFilter map[string]string

//=============================================================================================

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletFilter() *getWalletsFilter {
	return &getWalletsFilter{}
}

func (g *getWalletsFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getWalletsFilter) SetID(id string) *getWalletsFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getWalletsFilter) SetIDWithIN(id ...string) *getWalletsFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getWalletsFilter) SetIDWithNIN(id ...string) *getWalletsFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getWalletsFilter) SetLocked(locked bool) *getWalletsFilter {
	(*g)["filter[locked]"] = strconv.FormatBool(locked)
	return g
}

func (g *getWalletsFilter) SetName(name string) *getWalletsFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getWalletsFilter) SetOwnerID(ownerId string) *getWalletsFilter {
	(*g)["filter[ownerId]"] = ownerId
	return g
}

func (g *getWalletsFilter) SetReference(reference string) *getWalletsFilter {
	(*g)["filter[reference]"] = reference
	return g
}

func (g *getWalletsFilter) SetSortWithCreatedAt() *getWalletsFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletsFilter) SetSortWithCreatedAtReverse() *getWalletsFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗         █████╗  ██████╗████████╗██╗██╗   ██╗██╗████████╗██╗   ██╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝╚══██╔══╝██║██║   ██║██║╚══██╔══╝╚██╗ ██╔╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ███████║██║        ██║   ██║██║   ██║██║   ██║    ╚████╔╝         █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██║██║        ██║   ██║╚██╗ ██╔╝██║   ██║     ╚██╔╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║  ██║╚██████╗   ██║   ██║ ╚████╔╝ ██║   ██║      ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝  ╚═══╝  ╚═╝   ╚═╝      ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletActivtyWithIDFilter() *getWalletActivityWithIDFilter {
	return &getWalletActivityWithIDFilter{}
}

func (g *getWalletActivityWithIDFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletActivityWithIDFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getWalletActivityWithIDFilter) SetAssetCode(assetCode string) *getWalletActivityWithIDFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityWithIDFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityWithIDFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetCode(code string) *getWalletActivityWithIDFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getWalletActivityWithIDFilter) SetCodeWithIN(code ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetCodeWithNIN(code ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetStatus(status string) *getWalletActivityWithIDFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *getWalletActivityWithIDFilter) SetStatusWithIN(status ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetStatusWithNIN(status ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetType(getType string) *getWalletActivityWithIDFilter {
	(*g)["filter[type]"] = getType
	return g
}

func (g *getWalletActivityWithIDFilter) SetTypeWithIN(getType ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(getType, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetTypeWithNIN(getType ...string) *getWalletActivityWithIDFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(getType, ",")
	return g
}

func (g *getWalletActivityWithIDFilter) SetSortWithAmount() *getWalletActivityWithIDFilter {
	(*g)["sort"] = "amount"
	return g
}

func (g *getWalletActivityWithIDFilter) SetSortWithAmountReverse() *getWalletActivityWithIDFilter {
	(*g)["sort"] = "-amount"
	return g
}

func (g *getWalletActivityWithIDFilter) SetSortWithCreatedAt() *getWalletActivityWithIDFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletActivityWithIDFilter) SetSortWithCreatedAtReverse() *getWalletActivityWithIDFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getWalletActivityWithIDFilter) SetAvailableIncludeWithUser() *getWalletActivityWithIDFilter {
	(*g)["include"] = "users"
	return g
}

func (g *getWalletActivityWithIDFilter) SetAvailableIncludeWithWallets() *getWalletActivityWithIDFilter {
	(*g)["include"] = "wallets"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ██████╗  █████╗ ██╗      █████╗ ███╗   ██╗ ██████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██║     ██╔══██╗████╗  ██║██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██████╔╝███████║██║     ███████║██╔██╗ ██║██║     █████╗          █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██╗██╔══██║██║     ██╔══██║██║╚██╗██║██║     ██╔══╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██████╔╝██║  ██║███████╗██║  ██║██║ ╚████║╚██████╗███████╗        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletBalanceWithIDFilter() *getWalletBalanceWithIDFilter {
	return &getWalletBalanceWithIDFilter{}
}

func (g *getWalletBalanceWithIDFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletBalanceWithIDFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceWithIDFilter) SetAssetCode(assetCode string) *getWalletBalanceWithIDFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceWithIDFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletBalanceWithIDFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceWithIDFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletBalanceWithIDFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ███╗   ███╗ ██████╗ ██╗   ██╗███████╗███╗   ███╗███████╗███╗   ██╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ████╗ ████║██╔═══██╗██║   ██║██╔════╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██╔████╔██║██║   ██║██║   ██║█████╗  ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletMovementsFilter() *getWalletMovementsFilter {
	return &getWalletMovementsFilter{}
}

func (g *getWalletMovementsFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletMovementsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementsFilter) SetAssetCode(assetCode string) *getWalletMovementsFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementsFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletMovementsFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementsFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletMovementsFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletMovementsFilter) SetSortWithAmount() *getWalletMovementsFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletMovementsFilter) SetSortWithAmountReverse() *getWalletMovementsFilter {
	(*g)["amount"] = "-amount"
	return g
}

func (g *getWalletMovementsFilter) SetSortWithCreatedAt() *getWalletMovementsFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletMovementsFilter) SetSortWithCreatedAtReverse() *getWalletMovementsFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getWalletMovementsFilter) SetAvailableTransactionWithUser() *getWalletMovementsFilter {
	(*g)["include"] = "transaction"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║              ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║              ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║              ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝              ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝

*/

func GetWalletTransactionsFilter() *getWalletTransactionstFilter {
	return &getWalletTransactionstFilter{}
}

func (g *getWalletTransactionstFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletTransactionstFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getWalletTransactionstFilter) SetAssetCode(assetCode string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletTransactionstFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletTransactionstFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetCode(code string) *getWalletTransactionstFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getWalletTransactionstFilter) SetCodeWithIN(code ...string) *getWalletTransactionstFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetCodeWithNIN(code ...string) *getWalletTransactionstFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetStatus(status string) *getWalletTransactionstFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *getWalletTransactionstFilter) SetStatusWithIN(status ...string) *getWalletTransactionstFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetStatusWithNIN(status ...string) *getWalletTransactionstFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetType(getType string) *getWalletTransactionstFilter {
	(*g)["filter[type]"] = getType
	return g
}

func (g *getWalletTransactionstFilter) SetTypeWithIN(getType ...string) *getWalletTransactionstFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(getType, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetTypeWithNIN(getType ...string) *getWalletTransactionstFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(getType, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetSortWithAmount() *getWalletTransactionstFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletTransactionstFilter) SetSortWithAmountReverse() *getWalletTransactionstFilter {
	(*g)["amount"] = "-amount"
	return g
}

func (g *getWalletTransactionstFilter) SetSortWithCreatedAt() *getWalletTransactionstFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletTransactionstFilter) SetSortWithCreatedAtReverse() *getWalletTransactionstFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗         █████╗  ██████╗████████╗██╗██╗   ██╗██╗████████╗██╗   ██╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝╚══██╔══╝██║██║   ██║██║╚══██╔══╝╚██╗ ██╔╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ███████║██║        ██║   ██║██║   ██║██║   ██║    ╚████╔╝         █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██║██║        ██║   ██║╚██╗ ██╔╝██║   ██║     ╚██╔╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║  ██║╚██████╗   ██║   ██║ ╚████╔╝ ██║   ██║      ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝  ╚═══╝  ╚═╝   ╚═╝      ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletActivityFilter() *getWalletActivityFilter {
	return &getWalletActivityFilter{}
}

func (g *getWalletActivityFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletActivityFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCode(assetCode string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletActivityFilter) SetCode(code string) *getWalletActivityFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getWalletActivityFilter) SetCodeWithIN(code ...string) *getWalletActivityFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getWalletActivityFilter) SetCodeWithNIN(code ...string) *getWalletActivityFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

func (g *getWalletActivityFilter) SetStatus(status string) *getWalletActivityFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *getWalletActivityFilter) SetStatusWithIN(status ...string) *getWalletActivityFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *getWalletActivityFilter) SetStatusWithNIN(status ...string) *getWalletActivityFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *getWalletActivityFilter) SetType(getType string) *getWalletActivityFilter {
	(*g)["filter[type]"] = getType
	return g
}

func (g *getWalletActivityFilter) SetTypeWithIN(getType ...string) *getWalletActivityFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(getType, ",")
	return g
}

func (g *getWalletActivityFilter) SetTypeWithNIN(getType ...string) *getWalletActivityFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(getType, ",")
	return g
}

func (g *getWalletActivityFilter) SetSortWithAmount() *getWalletActivityFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletActivityFilter) SetSortWithAmountReverse() *getWalletActivityFilter {
	(*g)["amount"] = "-amount"
	return g
}

func (g *getWalletActivityFilter) SetSortWithCreatedAt() *getWalletActivityFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletActivityFilter) SetSortWithCreatedAtReverse() *getWalletActivityFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getWalletActivityFilter) SetAvailableIncludeWithUser() *getWalletActivityFilter {
	(*g)["include"] = "users"
	return g
}

func (g *getWalletActivityFilter) SetAvailableIncludeWithWallets() *getWalletActivityFilter {
	(*g)["include"] = "wallets"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ██████╗  █████╗ ██╗      █████╗ ███╗   ██╗ ██████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██║     ██╔══██╗████╗  ██║██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██████╔╝███████║██║     ███████║██╔██╗ ██║██║     █████╗          █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██╗██╔══██║██║     ██╔══██║██║╚██╗██║██║     ██╔══╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██████╔╝██║  ██║███████╗██║  ██║██║ ╚████║╚██████╗███████╗        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetWalletBalanceFilter is
func GetWalletBalanceFilter() *getWalletBalanceFilter {
	return &getWalletBalanceFilter{}
}

func (g *getWalletBalanceFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletBalanceFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceFilter) SetAssetCode(assetCode string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗          ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝          ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗        ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// UpdateWalletBody is
func UpdateWalletDataForBody() *UpdateWalletDataBody {
	return &UpdateWalletDataBody{}
}

func (u *UpdateWalletDataBody) SetOwnerID(ownerID string) *UpdateWalletDataBody {
	u.OwnerID = ownerID
	return u
}

func (u *UpdateWalletDataBody) SetLocked(locked bool) *UpdateWalletDataBody {
	u.Locked = locked
	return u
}

func (u *UpdateWalletDataBody) SetCustomData(customData interface{}) *UpdateWalletDataBody {
	u.CustomData = customData
	return u
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateWalletBody() *CreateWalletDataBody {
	return &CreateWalletDataBody{}
}

func (c *CreateWalletDataBody) SetName(name string) *CreateWalletDataBody {
	c.Name = name
	return c
}

func (c *CreateWalletDataBody) SetOwnerID(ownerID string) *CreateWalletDataBody {
	c.OwnerID = ownerID
	return c
}

func (c *CreateWalletDataBody) SetLocked(locked bool) *CreateWalletDataBody {
	c.Locked = locked
	return c
}

func (c *CreateWalletDataBody) SetCustomData(customData interface{}) *CreateWalletDataBody {
	c.CustomData = customData
	return c
}
