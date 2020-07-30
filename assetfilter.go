package slyk

import (
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
func (u *getassetFilter) SetPageSize(size int) *getassetFilter {
	(*u)["page[size]"] = strconv.Itoa(size)
	return u
}

// Defines the number of the page to retrieve. Default = 1
func (u *getassetFilter) SetPageNumber(number int) *getassetFilter {
	(*u)["page[number]"] = strconv.Itoa(number)
	return u
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

func (u *UpdateAssetDataBody) SetCustomData(customData struct{}) *UpdateAssetDataBody {
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
