package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getTaxRatesFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ████████╗ █████╗ ██╗  ██╗        ██████╗  █████╗ ████████╗███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ╚══██╔══╝██╔══██╗╚██╗██╔╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║              ██║   ███████║ ╚███╔╝         ██████╔╝███████║   ██║   █████╗  ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║              ██║   ██╔══██║ ██╔██╗         ██╔══██╗██╔══██║   ██║   ██╔══╝  ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║              ██║   ██║  ██║██╔╝ ██╗        ██║  ██║██║  ██║   ██║   ███████╗███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝              ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetTaxRatesFilter is
func GetTaxRatesFilter() *getTaxRatesFilter {
	return &getTaxRatesFilter{}
}

func (g *getTaxRatesFilter) SetGenericQueryParameter(key string, value interface{}) *getTaxRatesFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getTaxRatesFilter) SetID(id string) *getTaxRatesFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getTaxRatesFilter) SetIDWithIN(id ...string) *getTaxRatesFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getTaxRatesFilter) SetIDWithNIN(id ...string) *getTaxRatesFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getTaxRatesFilter) SetName(name string) *getTaxRatesFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getTaxRatesFilter) SetSortWithCreatedAt() *getTaxRatesFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getTaxRatesFilter) SetSortWithCreatedAtReverse() *getTaxRatesFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getTaxRatesFilter) SetSortWithName() *getTaxRatesFilter {
	(*g)["sort"] = "name"
	return g
}

func (g *getTaxRatesFilter) SetSortWithNameReverse() *getTaxRatesFilter {
	(*g)["sort"] = "-name"
	return g
}

func (g *getTaxRatesFilter) SetSortWithRate() *getTaxRatesFilter {
	(*g)["sort"] = "rate"
	return g
}

func (g *getTaxRatesFilter) SetSortWithRateReverse() *getTaxRatesFilter {
	(*g)["sort"] = "-rate"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getTaxRatesFilter) SetPageSize(size int) *getTaxRatesFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getTaxRatesFilter) SetPageNumber(number int) *getTaxRatesFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗        ████████╗ █████╗ ██╗  ██╗        ██████╗  █████╗ ████████╗███████╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗╚██╗██╔╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗             ██║   ███████║ ╚███╔╝         ██████╔╝███████║   ██║   █████╗          ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝             ██║   ██╔══██║ ██╔██╗         ██╔══██╗██╔══██║   ██║   ██╔══╝          ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗           ██║   ██║  ██║██╔╝ ██╗        ██║  ██║██║  ██║   ██║   ███████╗        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// UpdateTaxRateDataForBody is
func UpdateTaxRateDataForBody() *UpdateTaxRateDataBody {
	return &UpdateTaxRateDataBody{}
}

func (u *UpdateTaxRateDataBody) SetName(name string) *UpdateTaxRateDataBody {
	u.Name = name
	return u
}

func (u *UpdateTaxRateDataBody) SetRate(rate string) *UpdateTaxRateDataBody {
	u.Rate = rate
	return u
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ████████╗ █████╗ ██╗  ██╗        ██████╗  █████╗ ████████╗███████╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗╚██╗██╔╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗             ██║   ███████║ ╚███╔╝         ██████╔╝███████║   ██║   █████╗          ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝             ██║   ██╔══██║ ██╔██╗         ██╔══██╗██╔══██║   ██║   ██╔══╝          ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗           ██║   ██║  ██║██╔╝ ██╗        ██║  ██║██║  ██║   ██║   ███████╗        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CreateTaxRateDataForBody is
func CreateTaxRateDataForBody() *CreateTaxRateDataBody {
	return &CreateTaxRateDataBody{}
}

func (u *CreateTaxRateDataBody) SetName(name string) *CreateTaxRateDataBody {
	u.Name = name
	return u
}

func (u *CreateTaxRateDataBody) SetRate(rate string) *CreateTaxRateDataBody {
	u.Rate = rate
	return u
}
