package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getCategoriesFilter map[string]string

/*
 ██████╗ ███████╗████████╗         ██████╗ █████╗ ████████╗███████╗ ██████╗  ██████╗ ██████╗ ██╗███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔════╝██╔══██╗╚══██╔══╝██╔════╝██╔════╝ ██╔═══██╗██╔══██╗██║██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║     ███████║   ██║   █████╗  ██║  ███╗██║   ██║██████╔╝██║█████╗  ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║     ██╔══██║   ██║   ██╔══╝  ██║   ██║██║   ██║██╔══██╗██║██╔══╝  ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚██████╗██║  ██║   ██║   ███████╗╚██████╔╝╚██████╔╝██║  ██║██║███████╗███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetCategoriesFilter is
func GetCategoriesFilter() *getCategoriesFilter {
	return &getCategoriesFilter{}
}

func (g *getCategoriesFilter) SetGenericQueryParameter(key string, value interface{}) *getCategoriesFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getCategoriesFilter) SetAvailableProducts(availableProducts interface{}) *getCategoriesFilter {
	(*g)["availableProducts"] = fmt.Sprintf("%v", availableProducts)
	return g
}

func (g *getCategoriesFilter) SetDescription(description string) *getCategoriesFilter {
	(*g)["description"] = description
	return g
}

func (g *getassetFilter) SetID(id string) *getassetFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getassetFilter) SetIDWithIN(id ...string) *getassetFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getassetFilter) SetIDWithNIN(id ...string) *getassetFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getassetFilter) SetOrder(order string) *getassetFilter {
	(*g)["filter[order]"] = order
	return g
}

func (g *getassetFilter) SetOrderWithGTE(order string) *getassetFilter {
	(*g)["filter[order]"] = "gte:" + order
	return g
}

func (g *getassetFilter) SetOrderWithLTE(order string) *getassetFilter {
	(*g)["filter[getassetFilter]"] = "lte:" + order
	return g
}

func (g *getassetFilter) SetTitle(title string) *getassetFilter {
	(*g)["filter[title]"] = title
	return g
}

func (g *getCategoriesFilter) SetVisibleProducts(visibleProducts interface{}) *getCategoriesFilter {
	(*g)["visibleProducts"] = fmt.Sprintf("%v", visibleProducts)
	return g
}

func (g *getCategoriesFilter) SetSortWithCreatedAt() *getCategoriesFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getCategoriesFilter) SetSortWithCreatedAtReverse() *getCategoriesFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getCategoriesFilter) SetSortWithOrder() *getCategoriesFilter {
	(*g)["sort"] = "order"
	return g
}

func (g *getCategoriesFilter) SetSortWithOrderReverse() *getCategoriesFilter {
	(*g)["sort"] = "-order"
	return g
}

func (g *getCategoriesFilter) SetSortWithTitle() *getCategoriesFilter {
	(*g)["sort"] = "title"
	return g
}

func (g *getCategoriesFilter) SetSortWithTitleReverse() *getCategoriesFilter {
	(*g)["sort"] = "-title"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getCategoriesFilter) SetPageSize(size int) *getCategoriesFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getCategoriesFilter) SetPageNumber(number int) *getCategoriesFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗         ██████╗ █████╗ ████████╗███████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔════╝██╔══██╗╚══██╔══╝██╔════╝██╔════╝ ██╔═══██╗██╔══██╗╚██╗ ██╔╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗          ██║     ███████║   ██║   █████╗  ██║  ███╗██║   ██║██████╔╝ ╚████╔╝         ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝          ██║     ██╔══██║   ██║   ██╔══╝  ██║   ██║██║   ██║██╔══██╗  ╚██╔╝          ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗        ╚██████╗██║  ██║   ██║   ███████╗╚██████╔╝╚██████╔╝██║  ██║   ██║           ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func UpdateCategoryDataForBody() *UpdateCategoryDataBody {
	return &UpdateCategoryDataBody{}
}

func (u *UpdateCategoryDataBody) SetDescription(description string) *UpdateCategoryDataBody {
	u.Description = description
	return u
}

func (u *UpdateCategoryDataBody) SetImage(image string) *UpdateCategoryDataBody {
	u.Image = image
	return u
}

func (u *UpdateCategoryDataBody) SetTitle(title string) *UpdateCategoryDataBody {
	u.Title = title
	return u
}

func (u *UpdateCategoryDataBody) SetCustomData(customData string) *UpdateCategoryDataBody {
	u.CustomData = customData
	return u
}

func (u *UpdateCategoryDataBody) SetOrder(order string) *UpdateCategoryDataBody {
	u.Order = order
	return u
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗         ██████╗ █████╗ ████████╗███████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔════╝██╔══██╗╚══██╔══╝██╔════╝██╔════╝ ██╔═══██╗██╔══██╗╚██╗ ██╔╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║     ███████║   ██║   █████╗  ██║  ███╗██║   ██║██████╔╝ ╚████╔╝         ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║     ██╔══██║   ██║   ██╔══╝  ██║   ██║██║   ██║██╔══██╗  ╚██╔╝          ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚██████╗██║  ██║   ██║   ███████╗╚██████╔╝╚██████╔╝██║  ██║   ██║           ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

func CreateCategoryDataForBody() *CreateCategoryDataBody {
	return &CreateCategoryDataBody{}
}

func (u *CreateCategoryDataBody) SetDescription(description string) *CreateCategoryDataBody {
	u.Description = description
	return u
}

func (u *CreateCategoryDataBody) SetImage(image string) *CreateCategoryDataBody {
	u.Image = image
	return u
}

func (u *CreateCategoryDataBody) SetTitle(title string) *CreateCategoryDataBody {
	u.Title = title
	return u
}

func (u *CreateCategoryDataBody) SetCustomData(customData interface{}) *CreateCategoryDataBody {
	u.CustomData = customData
	return u
}

func (u *CreateCategoryDataBody) SetOrder(order string) *CreateCategoryDataBody {
	u.Order = order
	return u
}

/*
 ██████╗ █████╗ ████████╗███████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗        ██████╗ ███████╗ ██████╗ ██████╗ ██████╗ ███████╗██████╗         ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗╚══██╔══╝██╔════╝██╔════╝ ██╔═══██╗██╔══██╗╚██╗ ██╔╝        ██╔══██╗██╔════╝██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ███████║   ██║   █████╗  ██║  ███╗██║   ██║██████╔╝ ╚████╔╝         ██████╔╝█████╗  ██║   ██║██████╔╝██║  ██║█████╗  ██████╔╝        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██║   ██║   ██╔══╝  ██║   ██║██║   ██║██╔══██╗  ╚██╔╝          ██╔══██╗██╔══╝  ██║   ██║██╔══██╗██║  ██║██╔══╝  ██╔══██╗        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║   ██║   ███████╗╚██████╔╝╚██████╔╝██║  ██║   ██║           ██║  ██║███████╗╚██████╔╝██║  ██║██████╔╝███████╗██║  ██║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝           ╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CategoryReorderDataForBody is
func CategoryReorderDataForBody() *CategoryReorderDataBody {
	return &CategoryReorderDataBody{}
}

func (c *CategoryReorderDataBody) SetSubsequentID(subsequentID string) *CategoryReorderDataBody {
	c.SubsequentID = subsequentID
	return c
}
