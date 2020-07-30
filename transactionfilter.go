package slyk

import "strings"

type geTransactionstFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║              ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║              ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║              ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝              ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetTransactionsFilter is
func GetTransactionsFilter() *geTransactionstFilter {
	return &geTransactionstFilter{}
}

func (g *geTransactionstFilter) SetAssetCodeWithIN(assetCode ...string) *geTransactionstFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

func (g *geTransactionstFilter) SetAssetCodeWithNIN(assetCode ...string) *geTransactionstFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *geTransactionstFilter) SetSortWithAmount(amount string) *geTransactionstFilter {
	(*g)["amount"] = "amount"
	return g
}

func (g *geTransactionstFilter) SetSortWithAmountReverse(amount string) *geTransactionstFilter {
	(*g)["amount"] = "-amount"
	return g
}

func (g *geTransactionstFilter) SetSortWithCreatedAt() *geTransactionstFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *geTransactionstFilter) SetSortWithCreatedAtReverse() *geTransactionstFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *geTransactionstFilter) SetAvailableIncludeTransaction() *geTransactionstFilter {
	(*g)["include"] = "transaction"
	return g
}
