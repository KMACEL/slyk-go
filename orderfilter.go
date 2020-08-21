package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getOrdersFilter map[string]string
type getOrderLinesWithIDFilter map[string]string

/*
 ██████╗ ███████╗████████╗         ██████╗ ██████╗ ██████╗ ███████╗██████╗ ███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetOrdersFilter() *getOrdersFilter {
	return &getOrdersFilter{}
}

func (g *getOrdersFilter) SetGenericQueryParameter(key string, value interface{}) *getOrdersFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getOrdersFilter) SetAmount(amount string) *getOrdersFilter {
	(*g)["filter[amount]"] = amount
	return g
}

func (g *getOrdersFilter) SetAmountWithGTE(amount string) *getOrdersFilter {
	(*g)["filter[amount]"] = "gte:" + amount
	return g
}

func (g *getOrdersFilter) SetAmountWithLTE(amount string) *getOrdersFilter {
	(*g)["filter[amount]"] = "lte:" + amount
	return g
}

func (g *getOrdersFilter) SetAssetCode(assetCode string) *getOrdersFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getOrdersFilter) SetAssetCodeWithIN(assetCode ...string) *getOrdersFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getOrdersFilter) SetAssetCodeWithNIN(assetCode ...string) *getOrdersFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getOrdersFilter) SetBonus(bonus string) *getOrdersFilter {
	(*g)["filter[bonus]"] = bonus
	return g
}

func (g *getOrdersFilter) SetBonusWithGTE(bonus string) *getOrdersFilter {
	(*g)["filter[bonus]"] = "gte:" + bonus
	return g
}

func (g *getOrdersFilter) SetBonusWithLTE(bonus string) *getOrdersFilter {
	(*g)["filter[bonus]"] = "lte:" + bonus
	return g
}

func (g *getOrdersFilter) SetFulfilledAt(fulfilledAt string) *getOrdersFilter {
	(*g)["filter[fulfilledAt]"] = fulfilledAt
	return g
}

func (g *getOrdersFilter) SetFulfilledAtWithGTE(fulfilledAt string) *getOrdersFilter {
	(*g)["filter[fulfilledAt]"] = "gte:" + fulfilledAt
	return g
}

func (g *getOrdersFilter) SetFulfilledAtWithLTE(fulfilledAt string) *getOrdersFilter {
	(*g)["filter[fulfilledAt]"] = "lte:" + fulfilledAt
	return g
}

func (g *getOrdersFilter) SetHideDraftsWithLTE(hideDrafts interface{}) *getOrdersFilter {
	(*g)["filter[hideDrafts]"] = fmt.Sprintf("%v", hideDrafts)
	return g
}

func (g *getOrdersFilter) SetOrderStatus(orderStatus string) *getOrdersFilter {
	(*g)["filter[orderStatus]"] = orderStatus
	return g
}

func (g *getOrdersFilter) SetOrderStatuseWithIN(orderStatus ...string) *getOrdersFilter {
	(*g)["filter[orderStatus]"] = "in:" + strings.Join(orderStatus, ",")
	return g
}

func (g *getOrdersFilter) SetOrderStatusWithNIN(orderStatus ...string) *getOrdersFilter {
	(*g)["filter[orderStatus]"] = "nin:" + strings.Join(orderStatus, ",")
	return g
}

func (g *getOrdersFilter) SetPaidAmount(paidAmount string) *getOrdersFilter {
	(*g)["filter[paidAmount]"] = paidAmount
	return g
}

func (g *getOrdersFilter) SetPaidAmountWithGTE(paidAmount string) *getOrdersFilter {
	(*g)["filter[paidAmount]"] = "gte:" + paidAmount
	return g
}

func (g *getOrdersFilter) SetPaidAmountWithLTE(paidAmount string) *getOrdersFilter {
	(*g)["filter[paidAmount]"] = "lte:" + paidAmount
	return g
}

func (g *getOrdersFilter) SetPaidAt(paidAt string) *getOrdersFilter {
	(*g)["filter[paidAt]"] = paidAt
	return g
}

func (g *getOrdersFilter) SetPaidAtWithGTE(paidAt string) *getOrdersFilter {
	(*g)["filter[paidAt]"] = "gte:" + paidAt
	return g
}

func (g *getOrdersFilter) SetPaidAtWithLTE(paidAt string) *getOrdersFilter {
	(*g)["filter[paidAt]"] = "lte:" + paidAt
	return g
}

func (g *getOrdersFilter) SetPaymentStatus(paymentStatus string) *getOrdersFilter {
	(*g)["filter[paymentStatus]"] = paymentStatus
	return g
}

func (g *getOrdersFilter) SetPaymentStatusWithIN(paymentStatus ...string) *getOrdersFilter {
	(*g)["filter[paymentStatus]"] = "in:" + strings.Join(paymentStatus, ",")
	return g
}

func (g *getOrdersFilter) SetPaymentStatusWithNIN(paymentStatus ...string) *getOrdersFilter {
	(*g)["filter[paymentStatus]"] = "nin:" + strings.Join(paymentStatus, ",")
	return g
}

func (g *getOrdersFilter) SetReference(reference string) *getOrdersFilter {
	(*g)["filter[reference]"] = reference
	return g
}

func (g *getOrdersFilter) SetReferenceWithIN(reference ...string) *getOrdersFilter {
	(*g)["filter[reference]"] = "in:" + strings.Join(reference, ",")
	return g
}

func (g *getOrdersFilter) SetReferenceNIN(reference ...string) *getOrdersFilter {
	(*g)["filter[reference]"] = "nin:" + strings.Join(reference, ",")
	return g
}

func (g *getOrdersFilter) SetTrackingID(trackingID string) *getOrdersFilter {
	(*g)["filter[trackingId]"] = trackingID
	return g
}

func (g *getOrdersFilter) SetTrackingIDWithIN(trackingID ...string) *getOrdersFilter {
	(*g)["filter[trackingId]"] = "in:" + strings.Join(trackingID, ",")
	return g
}

func (g *getOrdersFilter) SetTrackingIDWithNIN(trackingID ...string) *getOrdersFilter {
	(*g)["filter[trackingId]"] = "nin:" + strings.Join(trackingID, ",")
	return g
}

func (g *getOrdersFilter) SetUserID(userID string) *getOrdersFilter {
	(*g)["filter[userId]"] = userID
	return g
}

func (g *getOrdersFilter) SetUserIDWithIN(userID ...string) *getOrdersFilter {
	(*g)["filter[userId]"] = "in:" + strings.Join(userID, ",")
	return g
}

func (g *getOrdersFilter) SetUserIDWithNIN(userID ...string) *getOrdersFilter {
	(*g)["filter[userId]"] = "nin:" + strings.Join(userID, ",")
	return g
}

func (g *getOrdersFilter) SetSortWithCreatedAt() *getOrdersFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getOrdersFilter) SetSortWithCreatedAtReverse() *getOrdersFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getOrdersFilter) SetSortWithCanceledAt() *getOrdersFilter {
	(*g)["sort"] = "canceledAt"
	return g
}

func (g *getOrdersFilter) SetSortWithCanceledAtReverse() *getOrdersFilter {
	(*g)["sort"] = "-canceledAt"
	return g
}

func (g *getOrdersFilter) SetSortWithFulfilledAt() *getOrdersFilter {
	(*g)["sort"] = "fulfilledAt"
	return g
}

func (g *getOrdersFilter) SetSortWithFulfilledAtReverse() *getOrdersFilter {
	(*g)["sort"] = "-fulfilledAt"
	return g
}

func (g *getOrdersFilter) SetSortWithPaidAt() *getOrdersFilter {
	(*g)["sort"] = "paidAt"
	return g
}

func (g *getOrdersFilter) SetSortWithPaidAtReverse() *getOrdersFilter {
	(*g)["sort"] = "-paidAt"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getOrdersFilter) SetPageSize(size int) *getOrdersFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getOrdersFilter) SetPageNumber(number int) *getOrdersFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
 ██████╗ ███████╗████████╗         ██████╗ ██████╗ ██████╗ ███████╗██████╗         ██╗     ██╗███╗   ██╗███████╗███████╗        ██╗    ██╗██╗████████╗██╗  ██╗        ██╗██████╗         ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██║     ██║████╗  ██║██╔════╝██╔════╝        ██║    ██║██║╚══██╔══╝██║  ██║        ██║██╔══██╗        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        ██║     ██║██╔██╗ ██║█████╗  ███████╗        ██║ █╗ ██║██║   ██║   ███████║        ██║██║  ██║        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██║     ██║██║╚██╗██║██╔══╝  ╚════██║        ██║███╗██║██║   ██║   ██╔══██║        ██║██║  ██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ███████╗██║██║ ╚████║███████╗███████║        ╚███╔███╔╝██║   ██║   ██║  ██║        ██║██████╔╝        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚══════╝╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝         ╚══╝╚══╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝╚═════╝         ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetOrderLinesWithIDFilter is
func GetOrderLinesWithIDFilter() *getOrderLinesWithIDFilter {
	return &getOrderLinesWithIDFilter{}
}

func (g *getOrderLinesWithIDFilter) SetGenericQueryParameter(key string, value interface{}) *getOrderLinesWithIDFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getOrderLinesWithIDFilter) SetAssetCode(assetCode string) *getOrderLinesWithIDFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getOrderLinesWithIDFilter) SetAssetCodeWithIN(assetCode ...string) *getOrderLinesWithIDFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getOrderLinesWithIDFilter) SetAssetCodeWithNIN(assetCode ...string) *getOrderLinesWithIDFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getOrderLinesWithIDFilter) SetFulfilledAt(fulfilledAt string) *getOrderLinesWithIDFilter {
	(*g)["filter[fulfilledAt]"] = fulfilledAt
	return g
}

func (g *getOrderLinesWithIDFilter) SetFulfilledAtWithGTE(fulfilledAt string) *getOrderLinesWithIDFilter {
	(*g)["filter[fulfilledAt]"] = "gte:" + fulfilledAt
	return g
}

func (g *getOrderLinesWithIDFilter) SetFulfilledAtWithLTE(fulfilledAt string) *getOrderLinesWithIDFilter {
	(*g)["filter[fulfilledAt]"] = "lte:" + fulfilledAt
	return g
}

func (g *getOrderLinesWithIDFilter) SetFulfilledQuantity(fulfilledQuantity string) *getOrderLinesWithIDFilter {
	(*g)["filter[fulfilledQuantity]"] = fulfilledQuantity
	return g
}

func (g *getOrderLinesWithIDFilter) SetFulfilledQuantityWithGTE(fulfilledQuantity string) *getOrderLinesWithIDFilter {
	(*g)["filter[fulfilledQuantity]"] = "gte:" + fulfilledQuantity
	return g
}

func (g *getOrderLinesWithIDFilter) SetFulfilledQuantityWithLTE(fulfilledQuantity string) *getOrderLinesWithIDFilter {
	(*g)["filter[fulfilledQuantity]"] = "lte:" + fulfilledQuantity
	return g
}

func (g *getOrderLinesWithIDFilter) SetQuantity(quantity string) *getOrderLinesWithIDFilter {
	(*g)["filter[quantity]"] = quantity
	return g
}

func (g *getOrderLinesWithIDFilter) SetQuantityWithGTE(quantity string) *getOrderLinesWithIDFilter {
	(*g)["filter[quantity]"] = "gte:" + quantity
	return g
}

func (g *getOrderLinesWithIDFilter) SetQuantityWithLTE(quantity string) *getOrderLinesWithIDFilter {
	(*g)["filter[quantity]"] = "lte:" + quantity
	return g
}

func (g *getOrderLinesWithIDFilter) SetStatus(status string) *getOrderLinesWithIDFilter {
	(*g)["filter[status]"] = status
	return g
}

func (g *getOrderLinesWithIDFilter) SetStatusWithIN(status ...string) *getOrderLinesWithIDFilter {
	(*g)["filter[status]"] = "in:" + strings.Join(status, ",")
	return g
}

func (g *getOrderLinesWithIDFilter) SetStatusWithNIN(status ...string) *getOrderLinesWithIDFilter {
	(*g)["filter[status]"] = "nin:" + strings.Join(status, ",")
	return g
}

func (g *getOrderLinesWithIDFilter) SetUnitPrice(unitPrice string) *getOrderLinesWithIDFilter {
	(*g)["filter[unitPrice]"] = unitPrice
	return g
}

func (g *getOrderLinesWithIDFilter) SetUnitPriceWithGTE(unitPrice string) *getOrderLinesWithIDFilter {
	(*g)["filter[unitPrice]"] = "gte:" + unitPrice
	return g
}

func (g *getOrderLinesWithIDFilter) SetUnitPriceWithLTE(unitPrice string) *getOrderLinesWithIDFilter {
	(*g)["filter[unitPrice]"] = "lte:" + unitPrice
	return g
}

func (g *getOrderLinesWithIDFilter) SetSortWithCreatedAt() *getOrderLinesWithIDFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getOrderLinesWithIDFilter) SetSortWithCreatedAtReverse() *getOrderLinesWithIDFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getOrderLinesWithIDFilter) SetSortWithFulfilledAt() *getOrderLinesWithIDFilter {
	(*g)["sort"] = "fulfilledAt"
	return g
}

func (g *getOrderLinesWithIDFilter) SetSortWithFulfilledAtReverse() *getOrderLinesWithIDFilter {
	(*g)["sort"] = "-fulfilledAt"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getOrderLinesWithIDFilter) SetPageSize(size int) *getOrderLinesWithIDFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getOrderLinesWithIDFilter) SetPageNumber(number int) *getOrderLinesWithIDFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗         ██████╗ ██████╗ ██████╗ ███████╗██████╗         ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateOrderDataForBody() *CreateOrderDataBody {
	return &CreateOrderDataBody{}
}

func (c *CreateOrderDataBody) SetChosenPaymentMethod(chosenPaymentMethod string) *CreateOrderDataBody {
	c.ChosenPaymentMethod = chosenPaymentMethod
	return c
}

func (c *CreateOrderDataBody) SetCustomData(customData interface{}) *CreateOrderDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateOrderDataBody) SetDeliveryMethod(deliveryMethod string) *CreateOrderDataBody {
	c.DeliveryMethod = deliveryMethod
	return c
}

func (c *CreateOrderDataBody) SetDryRun(dryRun bool) *CreateOrderDataBody {
	c.DryRun = dryRun
	return c
}

func (c *CreateOrderDataBody) SetLines(lines []LineForOrder) *CreateOrderDataBody {
	c.Lines = lines
	return c
}

func (c *CreateOrderDataBody) AppendLine(line LineForOrder) *CreateOrderDataBody {
	c.Lines = append(c.Lines, line)
	return c
}

func (c *CreateOrderDataBody) SetShippingAddressID(shippingAddressID string) *CreateOrderDataBody {
	c.ShippingAddressID = shippingAddressID
	return c
}

func (c *CreateOrderDataBody) SetUseBonus(useBonus bool) *CreateOrderDataBody {
	c.UseBonus = useBonus
	return c
}

func (c *CreateOrderDataBody) SetUserID(userID string) *CreateOrderDataBody {
	c.UserID = userID
	return c
}

func (c *CreateOrderDataBody) SetUserNotes(userNotes string) *CreateOrderDataBody {
	c.UserNotes = userNotes
	return c
}

func (c *CreateOrderDataBody) SetWalletID(walletID string) *CreateOrderDataBody {
	c.WalletID = walletID
	return c
}

/*
 ██████╗ ██████╗ ██████╗ ███████╗██████╗          ██████╗ █████╗ ███╗   ██╗ ██████╗███████╗██╗             ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██╔════╝██╔══██╗████╗  ██║██╔════╝██╔════╝██║             ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        ██║     ███████║██╔██╗ ██║██║     █████╗  ██║             ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██║     ██╔══██║██║╚██╗██║██║     ██╔══╝  ██║             ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ╚██████╗██║  ██║██║ ╚████║╚██████╗███████╗███████╗        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝         ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚══════╝╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func OrderCancelDataForBody() *OrderCancelDataBody {
	return &OrderCancelDataBody{}
}

func (o *OrderCancelDataBody) SetReason(reason string) *OrderCancelDataBody {
	o.Reason = reason
	return o
}

func (o *OrderCancelDataBody) SetRefundAmount(refundAmount string) *OrderCancelDataBody {
	o.RefundAmount = refundAmount
	return o
}

/*
 ██████╗ ██████╗ ██████╗ ███████╗██████╗         ███████╗██╗   ██╗██╗     ███████╗██╗██╗     ██╗             ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██╔════╝██║   ██║██║     ██╔════╝██║██║     ██║             ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        █████╗  ██║   ██║██║     █████╗  ██║██║     ██║             ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██╔══╝  ██║   ██║██║     ██╔══╝  ██║██║     ██║             ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ██║     ╚██████╔╝███████╗██║     ██║███████╗███████╗        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚══════╝╚═╝     ╚═╝╚══════╝╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// OrderFulfillDataForBody is
func OrderFulfillDataForBody() *OrderFulfillDataBody {
	return &OrderFulfillDataBody{}
}

func (o *OrderFulfillDataBody) SetTrackingID(trackingID string) *OrderFulfillDataBody {
	o.TrackingID = trackingID
	return o
}

func OrderPayDataForBody() *OrderPayDataBody {
	return &OrderPayDataBody{}
}

func (o *OrderPayDataBody) SetAmount(amount string) *OrderPayDataBody {
	o.Amount = amount
	return o
}

func (o *OrderPayDataBody) SetWalletID(walletID string) *OrderPayDataBody {
	o.WalletID = walletID
	return o
}

/*
 ██████╗ ██████╗ ██████╗ ███████╗██████╗         ██╗     ██╗███╗   ██╗███████╗        ███████╗██╗   ██╗██╗     ███████╗██╗██╗     ██╗             ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██║     ██║████╗  ██║██╔════╝        ██╔════╝██║   ██║██║     ██╔════╝██║██║     ██║             ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        ██║     ██║██╔██╗ ██║█████╗          █████╗  ██║   ██║██║     █████╗  ██║██║     ██║             ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██║     ██║██║╚██╗██║██╔══╝          ██╔══╝  ██║   ██║██║     ██╔══╝  ██║██║     ██║             ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ███████╗██║██║ ╚████║███████╗        ██║     ╚██████╔╝███████╗██║     ██║███████╗███████╗        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚══════╝╚═╝╚═╝  ╚═══╝╚══════╝        ╚═╝      ╚═════╝ ╚══════╝╚═╝     ╚═╝╚══════╝╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// OrderLineFulfillDataForBody is
func OrderLineFulfillDataForBody() *OrderLineFulfillDataBody {
	return &OrderLineFulfillDataBody{}
}

func (o *OrderLineFulfillDataBody) SetQuantity(quantity int) *OrderLineFulfillDataBody {
	o.Quantity = quantity
	return o
}

/*
 ██████╗ ██████╗ ██████╗ ███████╗██████╗         ██╗     ██╗███╗   ██╗███████╗        ██╗   ██╗███╗   ██╗        ███████╗██╗   ██╗██╗     ███████╗██╗██╗     ██╗             ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██║     ██║████╗  ██║██╔════╝        ██║   ██║████╗  ██║        ██╔════╝██║   ██║██║     ██╔════╝██║██║     ██║             ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        ██║     ██║██╔██╗ ██║█████╗          ██║   ██║██╔██╗ ██║        █████╗  ██║   ██║██║     █████╗  ██║██║     ██║             ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██║     ██║██║╚██╗██║██╔══╝          ██║   ██║██║╚██╗██║        ██╔══╝  ██║   ██║██║     ██╔══╝  ██║██║     ██║             ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ███████╗██║██║ ╚████║███████╗        ╚██████╔╝██║ ╚████║        ██║     ╚██████╔╝███████╗██║     ██║███████╗███████╗        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚══════╝╚═╝╚═╝  ╚═══╝╚══════╝         ╚═════╝ ╚═╝  ╚═══╝        ╚═╝      ╚═════╝ ╚══════╝╚═╝     ╚═╝╚══════╝╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// OrderLineUNFulfillDataForBody is
func OrderLineUNFulfillDataForBody() *OrderLineUNFulfillDataBody {
	return &OrderLineUNFulfillDataBody{}
}

func (o *OrderLineUNFulfillDataBody) SetQuantity(quantity int) *OrderLineUNFulfillDataBody {
	o.Quantity = quantity
	return o
}
