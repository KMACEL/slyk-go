package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (g *geTransactionstFilter) SetGenericQueryParameter(key string, value interface{}) *geTransactionstFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *geTransactionstFilter) SetAssetCode(assetCode string) *geTransactionstFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

func (g *geTransactionstFilter) SetAssetCodeWithIN(assetCode ...string) *geTransactionstFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

func (g *geTransactionstFilter) SetAssetCodeWithNIN(assetCode ...string) *geTransactionstFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *geTransactionstFilter) SetCode(code string) *geTransactionstFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *geTransactionstFilter) SetCodeWithIN(code ...string) *geTransactionstFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *geTransactionstFilter) SetCodeWithNIN(code ...string) *geTransactionstFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

// Format : 2019-07-21
func (g *geTransactionstFilter) SetCreated(date string) *geTransactionstFilter {
	(*g)["filter[createdAt]"] = date
	return g
}

// Format : 2019-07-21
func (g *geTransactionstFilter) SetCreatedAtWithGTE(date string) *geTransactionstFilter {
	(*g)["filter[createdAt]"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *geTransactionstFilter) SetCreatedAtWithLTE(date string) *geTransactionstFilter {
	(*g)["filter[createdAt]"] = "lte:" + date
	return g
}

func (g *geTransactionstFilter) SetDescription(description string) *geTransactionstFilter {
	(*g)["filter[description]"] = description
	return g
}

func (g *geTransactionstFilter) SetDestinationWalletID(destinationWalletID string) *geTransactionstFilter {
	(*g)["filter[destinationWalletId]"] = destinationWalletID
	return g
}

func (g *geTransactionstFilter) SetDestinationWalletIdWithIN(destinationWalletId ...string) *geTransactionstFilter {
	(*g)["filter[destinationWalletId]"] = "in:" + strings.Join(destinationWalletId, ",")
	return g
}

func (g *geTransactionstFilter) SetDestinationWalletIdWithNIN(destinationWalletId ...string) *geTransactionstFilter {
	(*g)["filter[destinationWalletId]"] = "nin:" + strings.Join(destinationWalletId, ",")
	return g
}

func (g *geTransactionstFilter) SetExternalId(externalId string) *geTransactionstFilter {
	(*g)["filter[externalId]"] = externalId
	return g
}

func (g *geTransactionstFilter) SetExternalIdWithIN(externalId ...string) *geTransactionstFilter {
	(*g)["filter[externalId]"] = "in:" + strings.Join(externalId, ",")
	return g
}

func (g *geTransactionstFilter) SetExternalIdWithNIN(externalId ...string) *geTransactionstFilter {
	(*g)["filter[externalId]"] = "nin:" + strings.Join(externalId, ",")
	return g
}

func (g *geTransactionstFilter) SetExternalReferenceWithNIN(externalReference string) *geTransactionstFilter {
	(*g)["filter[externalReference]"] = externalReference
	return g
}

func (g *geTransactionstFilter) SetID(id string) *geTransactionstFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *geTransactionstFilter) SetIDWithIN(id ...string) *geTransactionstFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *geTransactionstFilter) SetIDWithNIN(id ...string) *geTransactionstFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *geTransactionstFilter) SetOriginWalletID(originWalletId string) *geTransactionstFilter {
	(*g)["filter[originWalletId]"] = originWalletId
	return g
}

func (g *geTransactionstFilter) SetOriginWalletIDWithIN(originWalletId ...string) *geTransactionstFilter {
	(*g)["filter[originWalletId]"] = "in:" + strings.Join(originWalletId, ",")
	return g
}

func (g *geTransactionstFilter) SetOriginWalletIdWithNIN(originWalletId ...string) *geTransactionstFilter {
	(*g)["filter[originWalletId]"] = "nin:" + strings.Join(originWalletId, ",")
	return g
}

func (g *geTransactionstFilter) SetRSeference(reference string) *geTransactionstFilter {
	(*g)["filter[reference]"] = reference
	return g
}

func (g *geTransactionstFilter) SetStatus(status string) *geTransactionstFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *geTransactionstFilter) SetStatusWithIN(status ...string) *geTransactionstFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *geTransactionstFilter) SetStatusWithNIN(status ...string) *geTransactionstFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *geTransactionstFilter) SetType(typeParam string) *geTransactionstFilter {
	(*g)["filter[type]"] = typeParam
	return g
}

func (g *geTransactionstFilter) SetTypeWithIN(typeParam ...string) *geTransactionstFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(typeParam, ",")
	return g
}

func (g *geTransactionstFilter) SetTypeWithNIN(typeParam ...string) *geTransactionstFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(typeParam, ",")
	return g
}

func (g *geTransactionstFilter) SetWalletID(walletID string) *geTransactionstFilter {
	(*g)["filter[walletId]"] = walletID
	return g
}

func (g *geTransactionstFilter) SetWalletIdWithIN(walletId ...string) *geTransactionstFilter {
	(*g)["filter[walletId]"] = "in:" + strings.Join(walletId, ",")
	return g
}

func (g *geTransactionstFilter) SetWalletIdWithNIN(walletId ...string) *geTransactionstFilter {
	(*g)["filter[walletId]"] = "nin:" + strings.Join(walletId, ",")
	return g
}

func (g *geTransactionstFilter) SetSortWithAmount() *geTransactionstFilter {
	(*g)["sort"] = "amount"
	return g
}

func (g *geTransactionstFilter) SetSortWithAmountReverse() *geTransactionstFilter {
	(*g)["sort"] = "-amount"
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

// Defines the number of results per page. Default = 30.
func (u *geTransactionstFilter) SetPageSize(size int) *geTransactionstFilter {
	(*u)["page[size]"] = strconv.Itoa(size)
	return u
}

// Defines the number of the page to retrieve. Default = 1
func (u *geTransactionstFilter) SetPageNumber(number int) *geTransactionstFilter {
	(*u)["page[number]"] = strconv.Itoa(number)
	return u
}

/*
███████╗███████╗████████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ███████╗ █████╗ ██╗██╗             ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔════╝╚══██╔══╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔════╝██╔══██╗██║██║             ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
███████╗█████╗     ██║              ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║        █████╗  ███████║██║██║             ██████╔╝██║   ██║██║  ██║ ╚████╔╝
╚════██║██╔══╝     ██║              ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║        ██╔══╝  ██╔══██║██║██║             ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
███████║███████╗   ██║              ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║        ██║     ██║  ██║██║███████╗        ██████╔╝╚██████╔╝██████╔╝   ██║
╚══════╝╚══════╝   ╚═╝              ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝     ╚═╝  ╚═╝╚═╝╚══════╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝

*/

func TransactionFailDataForBody() *TransactionFailDataBody {
	return &TransactionFailDataBody{}
}

func (t *TransactionFailDataBody) SetReason(reason string) *TransactionFailDataBody {
	t.Reason = reason
	return t
}

/*
███████╗███████╗████████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ██████╗ ███████╗     ██╗███████╗ ██████╗████████╗        ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔════╝╚══██╔══╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔════╝     ██║██╔════╝██╔════╝╚══██╔══╝        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
███████╗█████╗     ██║              ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║        ██████╔╝█████╗       ██║█████╗  ██║        ██║           ██████╔╝██║   ██║██║  ██║ ╚████╔╝
╚════██║██╔══╝     ██║              ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║        ██╔══██╗██╔══╝  ██   ██║██╔══╝  ██║        ██║           ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
███████║███████╗   ██║              ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║        ██║  ██║███████╗╚█████╔╝███████╗╚██████╗   ██║           ██████╔╝╚██████╔╝██████╔╝   ██║
╚══════╝╚══════╝   ╚═╝              ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝  ╚═╝╚══════╝ ╚════╝ ╚══════╝ ╚═════╝   ╚═╝           ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝

*/

func TransactionRejectDataForBody() *TransactionRejectDataBody {
	return &TransactionRejectDataBody{}
}

func (t *TransactionRejectDataBody) SetReason(reason string) *TransactionRejectDataBody {
	t.Reason = reason
	return t
}

/*
 █████╗ ██████╗ ██████╗         ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ██████╗ ███████╗██████╗  ██████╗ ███████╗██╗████████╗        ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔══██╗██╔══██╗        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔════╝██║╚══██╔══╝        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
███████║██║  ██║██║  ██║           ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║        ██║  ██║█████╗  ██████╔╝██║   ██║███████╗██║   ██║           ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██╔══██║██║  ██║██║  ██║           ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║        ██║  ██║██╔══╝  ██╔═══╝ ██║   ██║╚════██║██║   ██║           ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
██║  ██║██████╔╝██████╔╝           ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║        ██████╔╝███████╗██║     ╚██████╔╝███████║██║   ██║           ██████╔╝╚██████╔╝██████╔╝   ██║
╚═╝  ╚═╝╚═════╝ ╚═════╝            ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═════╝ ╚══════╝╚═╝      ╚═════╝ ╚══════╝╚═╝   ╚═╝           ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateTransactionDepositDataForBody() *CreateTransactionDepositDataBody {
	return &CreateTransactionDepositDataBody{}
}

func (c *CreateTransactionDepositDataBody) SetAmount(amount string) *CreateTransactionDepositDataBody {
	c.Amount = amount
	return c
}

func (c *CreateTransactionDepositDataBody) SetAssetCode(assetCode string) *CreateTransactionDepositDataBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateTransactionDepositDataBody) SetCode(code string) *CreateTransactionDepositDataBody {
	c.Code = code
	return c
}

func (c *CreateTransactionDepositDataBody) SetCommit(commit bool) *CreateTransactionDepositDataBody {
	c.Commit = commit
	return c
}

func (c *CreateTransactionDepositDataBody) SetCustomData(customData interface{}) *CreateTransactionDepositDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateTransactionDepositDataBody) SetDescription(description string) *CreateTransactionDepositDataBody {
	c.Description = description
	return c
}

func (c *CreateTransactionDepositDataBody) SetDestinationAddress(destinationAddress string) *CreateTransactionDepositDataBody {
	c.DestinationAddress = destinationAddress
	return c
}

func (c *CreateTransactionDepositDataBody) SetDestinationWalletID(destinationWalletID string) *CreateTransactionDepositDataBody {
	c.DestinationWalletID = destinationWalletID
	return c
}

func (c *CreateTransactionDepositDataBody) SetExternalReference(externalReference string) *CreateTransactionDepositDataBody {
	c.ExternalReference = externalReference
	return c
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ██████╗  █████╗ ██╗   ██╗        ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔══██╗╚██╗ ██╔╝        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗             ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║        ██████╔╝███████║ ╚████╔╝         ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝             ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║        ██╔═══╝ ██╔══██║  ╚██╔╝          ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗           ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║        ██║     ██║  ██║   ██║           ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝     ╚═╝  ╚═╝   ╚═╝           ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateTransactionPayDataForBody() *CreateTransactionPayDataBody {
	return &CreateTransactionPayDataBody{}
}

func (c *CreateTransactionPayDataBody) SetAmount(amount string) *CreateTransactionPayDataBody {
	c.Amount = amount
	return c
}

func (c *CreateTransactionPayDataBody) SetAssetCode(assetCode string) *CreateTransactionPayDataBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateTransactionPayDataBody) SetCustomData(customData interface{}) *CreateTransactionPayDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateTransactionPayDataBody) SetDescription(description string) *CreateTransactionPayDataBody {
	c.Description = description
	return c
}

func (c *CreateTransactionPayDataBody) SetOriginWalletID(originWalletID string) *CreateTransactionPayDataBody {
	c.OriginWalletID = originWalletID
	return c
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗███████╗███████╗██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔════╝██╔════╝██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗             ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║           ██║   ██████╔╝███████║██╔██╗ ██║███████╗█████╗  █████╗  ██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝             ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║           ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══╝  ██╔══╝  ██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗           ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║           ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║     ███████╗██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝           ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateTransactionTransferDataForBody() *CreateTransactionTransferDataBody {
	return &CreateTransactionTransferDataBody{}
}

func (c *CreateTransactionTransferDataBody) SetAmount(amount string) *CreateTransactionTransferDataBody {
	c.Amount = amount
	return c
}

func (c *CreateTransactionTransferDataBody) SetAssetCode(assetCode string) *CreateTransactionTransferDataBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateTransactionTransferDataBody) SetCode(code string) *CreateTransactionTransferDataBody {
	c.Code = code
	return c
}

func (c *CreateTransactionTransferDataBody) SetCommit(commit bool) *CreateTransactionTransferDataBody {
	c.Commit = commit
	return c
}

func (c *CreateTransactionTransferDataBody) SetCustomData(customData interface{}) *CreateTransactionTransferDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateTransactionTransferDataBody) SetDescription(description string) *CreateTransactionTransferDataBody {
	c.Description = description
	return c
}

func (c *CreateTransactionTransferDataBody) SetDestinationAddress(destinationAddress string) *CreateTransactionTransferDataBody {
	c.DestinationAddress = destinationAddress
	return c
}

func (c *CreateTransactionTransferDataBody) SetDestinationWalletID(destinationWalletID string) *CreateTransactionTransferDataBody {
	c.DestinationWalletID = destinationWalletID
	return c
}

func (c *CreateTransactionTransferDataBody) SetExternalReference(externalReference string) *CreateTransactionTransferDataBody {
	c.ExternalReference = externalReference
	return c
}

func (c *CreateTransactionTransferDataBody) SetOriginWalletID(originWalletID string) *CreateTransactionTransferDataBody {
	c.OriginWalletID = originWalletID
	return c
}

func (c *CreateTransactionTransferDataBody) SetOriginAddress(originAddress string) *CreateTransactionTransferDataBody {
	c.OriginAddress = originAddress
	return c
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ████████╗██████╗  █████╗ ███╗   ██╗███████╗ █████╗  ██████╗████████╗██╗ ██████╗ ███╗   ██╗        ██╗    ██╗██╗████████╗██╗  ██╗██████╗ ██████╗  █████╗ ██╗    ██╗ █████╗ ██╗             ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██║    ██║██║╚══██╔══╝██║  ██║██╔══██╗██╔══██╗██╔══██╗██║    ██║██╔══██╗██║             ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗             ██║   ██████╔╝███████║██╔██╗ ██║███████╗███████║██║        ██║   ██║██║   ██║██╔██╗ ██║        ██║ █╗ ██║██║   ██║   ███████║██║  ██║██████╔╝███████║██║ █╗ ██║███████║██║             ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝             ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██╔══██║██║        ██║   ██║██║   ██║██║╚██╗██║        ██║███╗██║██║   ██║   ██╔══██║██║  ██║██╔══██╗██╔══██║██║███╗██║██╔══██║██║             ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗           ██║   ██║  ██║██║  ██║██║ ╚████║███████║██║  ██║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║        ╚███╔███╔╝██║   ██║   ██║  ██║██████╔╝██║  ██║██║  ██║╚███╔███╔╝██║  ██║███████╗        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝         ╚══╝╚══╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateTransactionWithdrawalDataForBody() *CreateTransactionWithdrawalDataBody {
	return &CreateTransactionWithdrawalDataBody{}
}

func (c *CreateTransactionWithdrawalDataBody) SetAmount(amount string) *CreateTransactionWithdrawalDataBody {
	c.Amount = amount
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetAssetCode(assetCode string) *CreateTransactionWithdrawalDataBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetCode(code string) *CreateTransactionWithdrawalDataBody {
	c.Code = code
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetCommit(commit bool) *CreateTransactionWithdrawalDataBody {
	c.Commit = commit
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetCustomData(customData interface{}) *CreateTransactionWithdrawalDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetDescription(description string) *CreateTransactionWithdrawalDataBody {
	c.Description = description
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetExternalReference(externalReference string) *CreateTransactionWithdrawalDataBody {
	c.ExternalReference = externalReference
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetOriginWalletID(originWalletID string) *CreateTransactionWithdrawalDataBody {
	c.OriginWalletID = originWalletID
	return c
}

func (c *CreateTransactionWithdrawalDataBody) SetOriginAddress(originAddress string) *CreateTransactionWithdrawalDataBody {
	c.OriginAddress = originAddress
	return c
}
