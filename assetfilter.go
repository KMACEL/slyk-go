package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getassetsFilter map[string]string

/*
 ██████╗ ███████╗████████╗         █████╗ ███████╗███████╗███████╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝██╔════╝██╔════╝╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ███████║███████╗███████╗█████╗     ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔══██║╚════██║╚════██║██╔══╝     ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║  ██║███████║███████║███████╗   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetAssetFilter is
func GetAssetsFilter() *getassetsFilter {
	return &getassetsFilter{}
}

func (g *getassetsFilter) SetGenericQueryParameter(key string, value interface{}) *getassetsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getassetsFilter) SetCode(code string) *getassetsFilter {
	(*g)["filter[code]"] = code
	return g
}

func (g *getassetsFilter) SetCodeWithIN(code ...string) *getassetsFilter {
	(*g)["filter[code]"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getassetsFilter) SetCodeWithNIN(code ...string) *getassetsFilter {
	(*g)["filter[code]"] = "nin:" + strings.Join(code, ",")
	return g
}

func (g *getassetsFilter) SetEnabled(enabled bool) *getassetsFilter {
	(*g)["filter[enabled]"] = strconv.FormatBool(enabled)
	return g
}

func (g *getassetsFilter) SetName(name string) *getassetsFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getassetsFilter) SetSystem(system bool) *getassetsFilter {
	(*g)["filter[system]"] = strconv.FormatBool(system)
	return g
}

func (g *getassetsFilter) SetType(getType string) *getassetsFilter {
	(*g)["filter[type]"] = getType
	return g
}

func (g *getassetsFilter) SetTypeWithIN(getType ...string) *getassetsFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(getType, ",")
	return g
}

func (g *getassetsFilter) SetTypeWithNIN(getType ...string) *getassetsFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(getType, ",")
	return g
}

func (g *getassetsFilter) SetSortWithCreatedAt() *getassetsFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getassetsFilter) SetSortWithCreatedAtReverse() *getassetsFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getassetsFilter) SetSortWithEnabled() *getassetsFilter {
	(*g)["sort"] = "enabled"
	return g
}

func (g *getassetsFilter) SetSortWithEnabledReverse() *getassetsFilter {
	(*g)["sort"] = "-enabled"
	return g
}

func (g *getassetsFilter) SetSortWithSystem() *getassetsFilter {
	(*g)["sort"] = "system"
	return g
}

func (g *getassetsFilter) SetSortWithSystemReverse() *getassetsFilter {
	(*g)["sort"] = "-system"
	return g
}

func (g *getassetsFilter) SetSortWithType() *getassetsFilter {
	(*g)["sort"] = "type"
	return g
}

func (g *getassetsFilter) SetSortWithTypeReverse() *getassetsFilter {
	(*g)["sort"] = "-type"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getassetsFilter) SetPageSize(size int) *getassetsFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getassetsFilter) SetPageNumber(number int) *getassetsFilter {
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
