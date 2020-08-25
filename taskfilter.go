package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getTasksFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ████████╗ █████╗ ███████╗██╗  ██╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║              ██║   ███████║███████╗█████╔╝ ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║              ██║   ██╔══██║╚════██║██╔═██╗ ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║              ██║   ██║  ██║███████║██║  ██╗███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝              ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetTasksFilter is
func GetTasksFilter() *getTasksFilter {
	return &getTasksFilter{}
}

func (g *getTasksFilter) SetGenericQueryParameter(key string, value interface{}) *getTasksFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getTasksFilter) SetEnabled(enabled bool) *getTasksFilter {
	(*g)["enabled"] = strconv.FormatBool(enabled)
	return g
}

func (g *getTasksFilter) SetFeatured(featured bool) *getTasksFilter {
	(*g)["featured"] = strconv.FormatBool(featured)
	return g
}

func (g *getTasksFilter) SetID(id string) *getTasksFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getTasksFilter) SetIDWithIN(id ...string) *getTasksFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getTasksFilter) SetIDWithNIN(id ...string) *getTasksFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getTasksFilter) SetName(name string) *getTasksFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getTasksFilter) SetNameWithIN(name ...string) *getTasksFilter {
	(*g)["filter[name]"] = "in:" + strings.Join(name, ",")
	return g
}

func (g *getTasksFilter) SetNameWithNIN(name ...string) *getTasksFilter {
	(*g)["filter[name]"] = "nin:" + strings.Join(name, ",")
	return g
}

func (g *getTasksFilter) SetOrderWithGTE(order string) *getTasksFilter {
	(*g)["filter[order]"] = "gte:" + order
	return g
}

func (g *getTasksFilter) SetOrdertWithLTE(order string) *getTasksFilter {
	(*g)["filter[order]"] = "lte:" + order
	return g
}

func (g *getTasksFilter) SetType(getType string) *getTasksFilter {
	(*g)["filter[type]"] = getType
	return g
}

func (g *getTasksFilter) SetTypeWithIN(getType ...string) *getTasksFilter {
	(*g)["filter[type]"] = "in:" + strings.Join(getType, ",")
	return g
}

func (g *getTasksFilter) SetTypeWithNIN(getType ...string) *getTasksFilter {
	(*g)["filter[type]"] = "nin:" + strings.Join(getType, ",")
	return g
}

func (g *getTasksFilter) SetSortWithAmount() *getTasksFilter {
	(*g)["sort"] = "amount"
	return g
}

func (g *getTasksFilter) SetSortWithAmountReverse() *getTasksFilter {
	(*g)["sort"] = "-amount"
	return g
}

func (g *getTasksFilter) SetSortWithCreatedAt() *getTasksFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getTasksFilter) SetSortWithCreatedAtReverse() *getTasksFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getTasksFilter) SetSortWithEnabled() *getTasksFilter {
	(*g)["sort"] = "enabled"
	return g
}

func (g *getTasksFilter) SetSortWithEnabledReverse() *getTasksFilter {
	(*g)["sort"] = "-enabled"
	return g
}

func (g *getTasksFilter) SetSortWithFeatured() *getTasksFilter {
	(*g)["sort"] = "featured"
	return g
}

func (g *getTasksFilter) SetSortWithFeaturedReverse() *getTasksFilter {
	(*g)["sort"] = "-featured"
	return g
}

func (g *getTasksFilter) SetSortWithName() *getTasksFilter {
	(*g)["sort"] = "name"
	return g
}

func (g *getTasksFilter) SetSortWithNameReverse() *getTasksFilter {
	(*g)["sort"] = "-name"
	return g
}

func (g *getTasksFilter) SetSortWithOrder() *getTasksFilter {
	(*g)["sort"] = "order"
	return g
}

func (g *getTasksFilter) SetSortWithOrderReverse() *getTasksFilter {
	(*g)["sort"] = "-order"
	return g
}

func (g *getTasksFilter) SetSortWithType() *getTasksFilter {
	(*g)["sort"] = "type"
	return g
}

func (g *getTasksFilter) SetSortWithTypeReverse() *getTasksFilter {
	(*g)["sort"] = "-type"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getTasksFilter) SetPageSize(size int) *getTasksFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getTasksFilter) SetPageNumber(number int) *getTasksFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}
