package slyk

import (
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

// asset btc,eth,ltc
func (g *getWalletActivityWithIDFilter) SetAssetCodeWithIN(assetCode string) *getWalletActivityWithIDFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityWithIDFilter) SetAssetCodeWithNIN(assetCode string) *getWalletActivityWithIDFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
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

func (g *getWalletActivityWithIDFilter) SetSortWithAmount(amount string) *getWalletActivityWithIDFilter {
	(*g)["filter[amount]"] = "amount"
	return g
}

func (g *getWalletActivityWithIDFilter) SetSortWithAmountReverse(amount string) *getWalletActivityWithIDFilter {
	(*g)["filter[amount]"] = "-amount"
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

// asset btc,eth,ltc
func (g *getWalletBalanceWithIDFilter) SetAssetCodeWithIN(assetCode string) *getWalletBalanceWithIDFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceWithIDFilter) SetAssetCodeWithNIN(assetCode string) *getWalletBalanceWithIDFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
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

func (g *getWalletMovementFilter) SetAssetCodeWithIN(assetCode string) *getWalletMovementFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementFilter) SetAssetCodeWithNIN(assetCode string) *getWalletMovementFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
	return g
}

func (g *getWalletMovementFilter) SetSortWithAmount(amount string) *getWalletMovementFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletMovementFilter) SetSortWithAmountReverse(amount string) *getWalletMovementFilter {
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

func (g *getWalletTransactionstFilter) SetAssetCodeWithIN(assetCode string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletTransactionstFilter) SetAssetCodeWithNIN(assetCode string) *getWalletTransactionstFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
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

func (g *getWalletTransactionstFilter) SetSortWithAmount(amount string) *getWalletTransactionstFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletTransactionstFilter) SetSortWithAmountReverse(amount string) *getWalletTransactionstFilter {
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

func (g *getWalletActivityFilter) SetAssetCodeWithIN(assetCode string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCodeWithNIN(assetCode string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
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

func (g *getWalletActivityFilter) SetSortWithAmount(amount string) *getWalletActivityFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *getWalletActivityFilter) SetSortWithAmountReverse(amount string) *getWalletActivityFilter {
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

func (g *getWalletBalanceFilter) SetAssetCodeWithIN(assetCode string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceFilter) SetAssetCodeWithNIN(assetCode string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
	return g
}
