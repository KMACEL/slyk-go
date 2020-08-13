package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getassetFilter map[string]string

/*
 ██████╗ ███████╗████████╗         █████╗ ███████╗███████╗███████╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝██╔════╝██╔════╝╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ███████║███████╗███████╗█████╗     ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔══██║╚════██║╚════██║██╔══╝     ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║  ██║███████║███████║███████╗   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetAssetFilter is
func GetAssetFilter() *getassetFilter {
	return &getassetFilter{}
}

func (g *getassetFilter) SetGenericQueryParameter(key string, value interface{}) *getassetFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getassetFilter) SetCode(code string) *getassetFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getassetFilter) SetCodeWithIN(code ...string) *getassetFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getassetFilter) SetCodeWithNIN(code ...string) *getassetFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

func (g *getassetFilter) SetEnabled(enabled bool) *getassetFilter {
	(*g)["filter[enabled]"] = strconv.FormatBool(enabled)
	return g
}

func (g *getassetFilter) SetName(name string) *getassetFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getassetFilter) SetSystem(system bool) *getassetFilter {
	(*g)["filter[system]"] = strconv.FormatBool(system)
	return g
}

func (g *getassetFilter) SetType(getType string) *getassetFilter {
	(*g)["filter[type]"] = getType
	return g
}

func (g *getassetFilter) SetTypeWithIN(getType ...string) *getassetFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(getType, ",")
	return g
}

func (g *getassetFilter) SetTypeWithNIN(getType ...string) *getassetFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(getType, ",")
	return g
}

func (g *getassetFilter) SetSortWithCreatedAt() *getassetFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getassetFilter) SetSortWithCreatedAtReverse() *getassetFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getassetFilter) SetSortWithEnabled() *getassetFilter {
	(*g)["sort"] = "enabled"
	return g
}

func (g *getassetFilter) SetSortWithEnabledReverse() *getassetFilter {
	(*g)["sort"] = "-enabled"
	return g
}

func (g *getassetFilter) SetSortWithSystem() *getassetFilter {
	(*g)["sort"] = "system"
	return g
}

func (g *getassetFilter) SetSortWithSystemReverse() *getassetFilter {
	(*g)["sort"] = "-system"
	return g
}

func (g *getassetFilter) SetSortWithType() *getassetFilter {
	(*g)["sort"] = "type"
	return g
}

func (g *getassetFilter) SetSortWithTypeReverse() *getassetFilter {
	(*g)["sort"] = "-type"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getassetFilter) SetPageSize(size int) *getassetFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getassetFilter) SetPageNumber(number int) *getassetFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗         █████╗ ███████╗███████╗███████╗████████╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔════╝██╔════╝██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗          ███████║███████╗███████╗█████╗     ██║           ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝          ██╔══██║╚════██║╚════██║██╔══╝     ██║           ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗        ██║  ██║███████║███████║███████╗   ██║           ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// GetAssetFilter is
func UpdateAssetDataForBody() *UpdateAssetDataBody {
	return &UpdateAssetDataBody{}
}

func (u *UpdateAssetDataBody) SetName(name string) *UpdateAssetDataBody {
	u.Name = name
	return u
}

func (u *UpdateAssetDataBody) SetDecimalPlaces(decimalPlaces int) *UpdateAssetDataBody {
	u.DecimalPlaces = decimalPlaces
	return u
}

func (u *UpdateAssetDataBody) SetContract(contract struct{}) *UpdateAssetDataBody {
	u.Contract = contract
	return u
}

func (u *UpdateAssetDataBody) SetCustomData(customData interface{}) *UpdateAssetDataBody {
	u.CustomData = customData
	return u
}

func (u *UpdateAssetDataBody) SetEnabled(enabled bool) *UpdateAssetDataBody {
	u.Enabled = enabled
	return u
}

func (u *UpdateAssetDataBody) SetLogo(logo string) *UpdateAssetDataBody {
	u.Logo = logo
	return u
}

func (u *UpdateAssetDataBody) SetSymbol(symbol string) *UpdateAssetDataBody {
	u.Symbol = symbol
	return u
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗         █████╗ ███████╗███████╗███████╗████████╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔════╝██╔════╝██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ███████║███████╗███████╗█████╗     ██║           ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██╔══██║╚════██║╚════██║██╔══╝     ██║           ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ██║  ██║███████║███████║███████╗   ██║           ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CreateAssetDataForBody
func CreateAssetDataForBody() *CreateAssetDataBody {
	return &CreateAssetDataBody{}
}

func (c *CreateAssetDataBody) SetCode(code string) *CreateAssetDataBody {
	c.Code = code
	return c
}

func (c *CreateAssetDataBody) SetName(name string) *CreateAssetDataBody {
	c.Name = name
	return c
}

func (c *CreateAssetDataBody) SetContract(contract struct{}) *CreateAssetDataBody {
	c.Contract = contract
	return c
}

func (c *CreateAssetDataBody) SetCustomData(customData interface{}) *CreateAssetDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateAssetDataBody) SetDecimalPlaces(decimalPlaces int) *CreateAssetDataBody {
	c.DecimalPlaces = decimalPlaces
	return c
}

func (c *CreateAssetDataBody) SetEnabled(enabled bool) *CreateAssetDataBody {
	c.Enabled = enabled
	return c
}

func (c *CreateAssetDataBody) SetSymbol(sysmbol string) *CreateAssetDataBody {
	c.Symbol = sysmbol
	return c
}

func (c *CreateAssetDataBody) SetType(typeParam string) *CreateAssetDataBody {
	c.Type = typeParam
	return c
}
