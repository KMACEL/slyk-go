package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getOrdersFilter map[string]string

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
