package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getWalletFilter map[string]string
type getWalletActivityWithIDFilter map[string]string
type getWalletBalanceWithIDFilter map[string]string
type getWalletMovementFilter map[string]string
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

func GetWalletFilter() *getWalletFilter {
	return &getWalletFilter{}
}

func (g *getWalletFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getWalletFilter) SetIDWithIN(id ...string) *getWalletFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getWalletFilter) SetIDWithNIN(id ...string) *getWalletFilter {
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

func (g *getWalletActivityWithIDFilter) SetCodeWithIN(code string) *getWalletActivityWithIDFilter {
	(*g)["filter[code]"] = "in:" + code
	return g
}

func (g *getWalletActivityWithIDFilter) SetCodeWithNIN(code string) *getWalletActivityWithIDFilter {
	(*g)["filter[code]"] = "nin:" + code
	return g
}

func (g *getWalletActivityWithIDFilter) SetStatusWithIN(status string) *getWalletActivityWithIDFilter {
	(*g)["filter[status]"] = "in:" + status
	return g
}

func (g *getWalletActivityWithIDFilter) SetStatusWithNIN(status string) *getWalletActivityWithIDFilter {
	(*g)["filter[status]"] = "nin:" + status
	return g
}

func (g *getWalletActivityWithIDFilter) SetTypeWithIN(getType string) *getWalletActivityWithIDFilter {
	(*g)["filter[type]"] = "in:" + getType
	return g
}

func (g *getWalletActivityWithIDFilter) SetTypeWithNIN(getType string) *getWalletActivityWithIDFilter {
	(*g)["filter[type]"] = "nin:" + getType
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

func GetWalletMovementFilter() *getWalletMovementFilter {
	return &getWalletMovementFilter{}
}

func (g *getWalletMovementFilter) SetGenericQueryParameter(key string, value interface{}) *getWalletMovementFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletMovementFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletMovementFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletMovementFilter) SetSortWithAmount() *getWalletMovementFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletMovementFilter) SetSortWithAmountReverse() *getWalletMovementFilter {
	(*g)["amount"] = "-amount"
	return g
}

func (g *getWalletMovementFilter) SetSortWithCreatedAt() *getWalletMovementFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletMovementFilter) SetSortWithCreatedAtReverse() *getWalletMovementFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getWalletMovementFilter) SetAvailableTransactionWithUser() *getWalletMovementFilter {
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
func (g *getWalletTransactionstFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletTransactionstFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletTransactionstFilter) SetCodeWithIN(code string) *getWalletTransactionstFilter {
	(*g)["filter[code]"] = "in:" + code
	return g
}

func (g *getWalletTransactionstFilter) SetCodeWithNIN(code string) *getWalletTransactionstFilter {
	(*g)["filter[code]"] = "nin:" + code
	return g
}

func (g *getWalletTransactionstFilter) SetStatusWithIN(status string) *getWalletTransactionstFilter {
	(*g)["filter[status]"] = "in:" + status
	return g
}

func (g *getWalletTransactionstFilter) SetStatusWithNIN(status string) *getWalletTransactionstFilter {
	(*g)["filter[status]"] = "nin:" + status
	return g
}

func (g *getWalletTransactionstFilter) SetTypeWithIN(getType string) *getWalletTransactionstFilter {
	(*g)["filter[type]"] = "in:" + getType
	return g
}

func (g *getWalletTransactionstFilter) SetTypeWithNIN(getType string) *getWalletTransactionstFilter {
	(*g)["filter[type]"] = "nin:" + getType
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
func (g *getWalletActivityFilter) SetAssetCodeWithIN(assetCode ...string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCodeWithNIN(assetCode ...string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getWalletActivityFilter) SetCodeWithIN(code string) *getWalletActivityFilter {
	(*g)["filter[code]"] = "in:" + code
	return g
}

func (g *getWalletActivityFilter) SetCodeWithNIN(code string) *getWalletActivityFilter {
	(*g)["filter[code]"] = "nin:" + code
	return g
}

func (g *getWalletActivityFilter) SetStatusWithIN(status string) *getWalletActivityFilter {
	(*g)["filter[status]"] = "in:" + status
	return g
}

func (g *getWalletActivityFilter) SetStatusWithNIN(status string) *getWalletActivityFilter {
	(*g)["filter[status]"] = "nin:" + status
	return g
}

func (g *getWalletActivityFilter) SetTypeWithIN(getType string) *getWalletActivityFilter {
	(*g)["filter[type]"] = "in:" + getType
	return g
}

func (g *getWalletActivityFilter) SetTypeWithNIN(getType string) *getWalletActivityFilter {
	(*g)["filter[type]"] = "nin:" + getType
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
func UpdateWalletBody() *UpdateWalletData {
	return &UpdateWalletData{}
}

func (u *UpdateWalletData) SetOwnerID(ownerID string) *UpdateWalletData {
	u.OwnerID = ownerID
	return u
}

func (u *UpdateWalletData) SetLocked(locked bool) *UpdateWalletData {
	u.Locked = locked
	return u
}

func (u *UpdateWalletData) SetCustomData(customData interface{}) *UpdateWalletData {
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

func CreateWalletBody() *CreateWalletData {
	return &CreateWalletData{}
}

func (c *CreateWalletData) SetName(name string) *CreateWalletData {
	c.Name = name
	return c
}

func (c *CreateWalletData) SetOwnerID(ownerID string) *CreateWalletData {
	c.OwnerID = ownerID
	return c
}

func (c *CreateWalletData) SetLocked(locked bool) *CreateWalletData {
	c.Locked = locked
	return c
}

// TODO Çalışmıyor bakılacak
func (c *CreateWalletData) SetCustomData(customData interface{}) *CreateWalletData {
	c.CustomData = customData
	return c
}
