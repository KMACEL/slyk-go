package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getRateFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██████╗  █████╗ ████████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██████╔╝███████║   ██║   █████╗          █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔══██╗██╔══██║   ██║   ██╔══╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║  ██║██║  ██║   ██║   ███████╗        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetRateFilter is
func GetRateFilter() *getRateFilter {
	return &getRateFilter{}
}

func (g *getRateFilter) SetGenericQueryParameter(key string, value interface{}) *getRateFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getRateFilter) SetBaseAssetCode(baseAssetCode string) *getRateFilter {
	(*g)["baseAssetCode"] = baseAssetCode
	return g
}

func (g *getRateFilter) SetBaseAssetCodeWithIN(baseAssetCode ...string) *getRateFilter {
	(*g)["baseAssetCode"] = "in:" + strings.Join(baseAssetCode, ",")
	return g
}

func (g *getRateFilter) SetBaseAssetCodeWithNIN(baseAssetCode ...string) *getRateFilter {
	(*g)["baseAssetCode"] = "nin:" + strings.Join(baseAssetCode, ",")
	return g
}

func (g *getRateFilter) SetQuoteAssetCode(quoteAssetCode string) *getRateFilter {
	(*g)["quoteAssetCode"] = quoteAssetCode
	return g
}

func (g *getRateFilter) SetQuoteAssetCodeWithIN(quoteAssetCode ...string) *getRateFilter {
	(*g)["quoteAssetCode"] = "in:" + strings.Join(quoteAssetCode, ",")
	return g
}

func (g *getRateFilter) SetQuoteAssetCodeWithNIN(quoteAssetCode ...string) *getRateFilter {
	(*g)["quoteAssetCode"] = "nin:" + strings.Join(quoteAssetCode, ",")
	return g
}

func (g *getRateFilter) SetRateWithGT(rate string) *getRateFilter {
	(*g)["rate"] = "gt:" + rate
	return g
}

func (g *getRateFilter) SetRateWithGTE(rate string) *getRateFilter {
	(*g)["rate"] = "gte:" + rate
	return g
}

func (g *getRateFilter) SetRateWithLT(rate string) *getRateFilter {
	(*g)["rate"] = "lt:" + rate
	return g
}

func (g *getRateFilter) SetRateWithLTE(rate string) *getRateFilter {
	(*g)["rate"] = "lte:" + rate
	return g
}

func (g *getRateFilter) SetSortWithCreatedAt() *getRateFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getRateFilter) SetSortWithCreatedAtReverse() *getRateFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getRateFilter) SetSortWithUpdatedAt() *getRateFilter {
	(*g)["sort"] = "updatedAt"
	return g
}

func (g *getRateFilter) SetSortWithUpdatedAtReverse() *getRateFilter {
	(*g)["sort"] = "-updatedAt"
	return g
}

func (g *getRateFilter) SetSortWitRate() *getRateFilter {
	(*g)["sort"] = "rate"
	return g
}

func (g *getRateFilter) SetSortWitRateReverse() *getRateFilter {
	(*g)["sort"] = "-rate"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getRateFilter) SetPageSize(size int) *getRateFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getRateFilter) SetPageNumber(number int) *getRateFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗    ██████╗  █████╗ ████████╗███████╗    ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝    ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝    ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗      ██████╔╝███████║   ██║   █████╗      ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝      ██╔══██╗██╔══██║   ██║   ██╔══╝      ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗    ██║  ██║██║  ██║   ██║   ███████╗    ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝    ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝    ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func UpdateRateBody() *UpdateRateBodyData {
	return &UpdateRateBodyData{}
}

func (c *UpdateRateBodyData) SetRate(rate string) *UpdateRateBodyData {
	c.Rate = rate
	return c
}

func (c *UpdateRateBodyData) SetCustomData(customData interface{}) *UpdateRateBodyData {
	c.CustomData = customData
	return c
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗    ██████╗  █████╗ ████████╗███████╗    ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝    ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝    ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗      ██████╔╝███████║   ██║   █████╗      ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝      ██╔══██╗██╔══██║   ██║   ██╔══╝      ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗    ██║  ██║██║  ██║   ██║   ███████╗    ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝    ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝    ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateRateBody() *CreateRateBodyData {
	return &CreateRateBodyData{}
}

func (c *CreateRateBodyData) SetBaseAssetCode(baseAssetCode string) *CreateRateBodyData {
	c.BaseAssetCode = baseAssetCode
	return c
}

func (c *CreateRateBodyData) SetQuoteAssetCode(quoteAssetCode string) *CreateRateBodyData {
	c.QuoteAssetCode = quoteAssetCode
	return c
}

func (c *CreateRateBodyData) SetRate(rate string) *CreateRateBodyData {
	c.Rate = rate
	return c
}

func (c *CreateRateBodyData) SetCustomData(customData interface{}) *CreateRateBodyData {
	c.CustomData = customData
	return c
}
