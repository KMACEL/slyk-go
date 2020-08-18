package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getRatesFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██████╗  █████╗ ████████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██████╔╝███████║   ██║   █████╗          █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔══██╗██╔══██║   ██║   ██╔══╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║  ██║██║  ██║   ██║   ███████╗        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetRateFilter is
func GetRatesFilter() *getRatesFilter {
	return &getRatesFilter{}
}

func (g *getRatesFilter) SetGenericQueryParameter(key string, value interface{}) *getRatesFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getRatesFilter) SetBaseAssetCode(baseAssetCode string) *getRatesFilter {
	(*g)["baseAssetCode"] = baseAssetCode
	return g
}

func (g *getRatesFilter) SetBaseAssetCodeWithIN(baseAssetCode ...string) *getRatesFilter {
	(*g)["baseAssetCode"] = "in:" + strings.Join(baseAssetCode, ",")
	return g
}

func (g *getRatesFilter) SetBaseAssetCodeWithNIN(baseAssetCode ...string) *getRatesFilter {
	(*g)["baseAssetCode"] = "nin:" + strings.Join(baseAssetCode, ",")
	return g
}

func (g *getRatesFilter) SetQuoteAssetCode(quoteAssetCode string) *getRatesFilter {
	(*g)["quoteAssetCode"] = quoteAssetCode
	return g
}

func (g *getRatesFilter) SetQuoteAssetCodeWithIN(quoteAssetCode ...string) *getRatesFilter {
	(*g)["quoteAssetCode"] = "in:" + strings.Join(quoteAssetCode, ",")
	return g
}

func (g *getRatesFilter) SetQuoteAssetCodeWithNIN(quoteAssetCode ...string) *getRatesFilter {
	(*g)["quoteAssetCode"] = "nin:" + strings.Join(quoteAssetCode, ",")
	return g
}

func (g *getRatesFilter) SetRateWithGT(rate string) *getRatesFilter {
	(*g)["rate"] = "gt:" + rate
	return g
}

func (g *getRatesFilter) SetRateWithGTE(rate string) *getRatesFilter {
	(*g)["rate"] = "gte:" + rate
	return g
}

func (g *getRatesFilter) SetRateWithLT(rate string) *getRatesFilter {
	(*g)["rate"] = "lt:" + rate
	return g
}

func (g *getRatesFilter) SetRateWithLTE(rate string) *getRatesFilter {
	(*g)["rate"] = "lte:" + rate
	return g
}

func (g *getRatesFilter) SetSortWithCreatedAt() *getRatesFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getRatesFilter) SetSortWithCreatedAtReverse() *getRatesFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getRatesFilter) SetSortWithUpdatedAt() *getRatesFilter {
	(*g)["sort"] = "updatedAt"
	return g
}

func (g *getRatesFilter) SetSortWithUpdatedAtReverse() *getRatesFilter {
	(*g)["sort"] = "-updatedAt"
	return g
}

func (g *getRatesFilter) SetSortWitRate() *getRatesFilter {
	(*g)["sort"] = "rate"
	return g
}

func (g *getRatesFilter) SetSortWitRateReverse() *getRatesFilter {
	(*g)["sort"] = "-rate"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getRatesFilter) SetPageSize(size int) *getRatesFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getRatesFilter) SetPageNumber(number int) *getRatesFilter {
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

func UpdateRateDataForBody() *UpdateRateDataBody {
	return &UpdateRateDataBody{}
}

func (c *UpdateRateDataBody) SetRate(rate string) *UpdateRateDataBody {
	c.Rate = rate
	return c
}

func (c *UpdateRateDataBody) SetCustomData(customData interface{}) *UpdateRateDataBody {
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

func CreateRateDataForBody() *CreateRateDataBody {
	return &CreateRateDataBody{}
}

func (c *CreateRateDataBody) SetBaseAssetCode(baseAssetCode string) *CreateRateDataBody {
	c.BaseAssetCode = baseAssetCode
	return c
}

func (c *CreateRateDataBody) SetQuoteAssetCode(quoteAssetCode string) *CreateRateDataBody {
	c.QuoteAssetCode = quoteAssetCode
	return c
}

func (c *CreateRateDataBody) SetRate(rate string) *CreateRateDataBody {
	c.Rate = rate
	return c
}

func (c *CreateRateDataBody) SetCustomData(customData interface{}) *CreateRateDataBody {
	c.CustomData = customData
	return c
}
