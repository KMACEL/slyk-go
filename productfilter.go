package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getProductsFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗ ██████╗████████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██║   ██║██╔════╝╚══██╔══╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██████╔╝██████╔╝██║   ██║██║  ██║██║   ██║██║        ██║   ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔═══╝ ██╔══██╗██║   ██║██║  ██║██║   ██║██║        ██║   ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║     ██║  ██║╚██████╔╝██████╔╝╚██████╔╝╚██████╗   ██║   ███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetProductsFilter is
func GetProductsFilter() *getProductsFilter {
	return &getProductsFilter{}
}

func (g *getProductsFilter) SetGenericQueryParameter(key string, value interface{}) *getProductsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getProductsFilter) SetAvailable(available string) *getProductsFilter {
	(*g)["filter[available]"] = available
	return g
}

func (g *getProductsFilter) SetCategoryID(categoryID string) *getProductsFilter {
	(*g)["filter[categoryId]"] = categoryID
	return g
}

func (g *getProductsFilter) SetCategoryIDWithIN(categoryID ...string) *getProductsFilter {
	(*g)["filter[categoryId]"] = "in:" + strings.Join(categoryID, ",")
	return g
}

func (g *getProductsFilter) SetCategoryIDWithNIN(categoryID ...string) *getProductsFilter {
	(*g)["filter[categoryId]"] = "nin:" + strings.Join(categoryID, ",")
	return g
}

func (g *getProductsFilter) SetDescription(description string) *getProductsFilter {
	(*g)["filter[description]"] = description
	return g
}

func (g *getProductsFilter) SetFeatured(featured bool) *getProductsFilter {
	(*g)["filter[featured]"] = strconv.FormatBool(featured)
	return g
}

func (g *getProductsFilter) SetID(id string) *getProductsFilter {
	(*g)["filter[id]"] = id
	return g
}

func (g *getProductsFilter) SetIDWithIN(id ...string) *getProductsFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getProductsFilter) SetIDWithNIN(id ...string) *getProductsFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getProductsFilter) SetName(name string) *getProductsFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getProductsFilter) SetOrderWithGTE(order string) *getProductsFilter {
	(*g)["filter[order]"] = "gte:" + order
	return g
}

func (g *getProductsFilter) SetOrdertWithLTE(order string) *getProductsFilter {
	(*g)["filter[order]"] = "lte:" + order
	return g
}

func (g *getProductsFilter) SetRequiresIdentity(requiresIdentity string) *getProductsFilter {
	(*g)["filter[requiresIdentity]"] = requiresIdentity
	return g
}

func (g *getProductsFilter) SetTypeCode(typeCode string) *getProductsFilter {
	(*g)["filter[typeCode]"] = typeCode
	return g
}

func (g *getProductsFilter) SetTypeCodeWithIN(typeCode ...string) *getProductsFilter {
	(*g)["filter[typeCode]"] = "in:" + strings.Join(typeCode, ",")
	return g
}

func (g *getProductsFilter) SetTypeCodeWithNIN(typeCode ...string) *getProductsFilter {
	(*g)["filter[typeCode]"] = "nin:" + strings.Join(typeCode, ",")
	return g
}

func (g *getProductsFilter) SetVisible(visible bool) *getProductsFilter {
	(*g)["filter[visible]"] = strconv.FormatBool(visible)
	return g
}

func (g *getProductsFilter) SetSortWithCreatedAt() *getProductsFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getProductsFilter) SetSortWithCreatedAtReverse() *getProductsFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getProductsFilter) SetSortWithOrder() *getProductsFilter {
	(*g)["sort"] = "order"
	return g
}

func (g *getProductsFilter) SetSortWithOrderReverse() *getProductsFilter {
	(*g)["sort"] = "-order"
	return g
}

func (g *getProductsFilter) SetSortWithFeatured() *getProductsFilter {
	(*g)["sort"] = "featured"
	return g
}

func (g *getProductsFilter) SetSortWithFeaturedReverse() *getProductsFilter {
	(*g)["sort"] = "-featured"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getProductsFilter) SetPageSize(size int) *getProductsFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getProductsFilter) SetPageNumber(number int) *getProductsFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ██████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗ ██████╗████████╗███████╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██║   ██║██╔════╝╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██████╔╝██████╔╝██║   ██║██║  ██║██║   ██║██║        ██║   ███████╗        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██╔═══╝ ██╔══██╗██║   ██║██║  ██║██║   ██║██║        ██║   ╚════██║        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ██║     ██║  ██║╚██████╔╝██████╔╝╚██████╔╝╚██████╗   ██║   ███████║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝   ╚═╝   ╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CreateProductsDataForBody is
func CreateProductsDataForBody() *CreateProductsDataBody {
	return &CreateProductsDataBody{}
}

func (c *CreateProductsDataBody) SetAllowChoosingQuantity(allowChoosingQuantity bool) *CreateProductsDataBody {
	c.AllowChoosingQuantity = allowChoosingQuantity
	return c
}

func (c *CreateProductsDataBody) SetAssetCode(assetCode string) *CreateProductsDataBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateProductsDataBody) SetAvailable(available bool) *CreateProductsDataBody {
	c.Available = available
	return c
}

func (c *CreateProductsDataBody) SetBonus(bonus string) *CreateProductsDataBody {
	c.Bonus = bonus
	return c
}

func (c *CreateProductsDataBody) SetButtonLabel(buttonLabel string) *CreateProductsDataBody {
	c.ButtonLabel = buttonLabel
	return c
}

func (c *CreateProductsDataBody) SetCategoryID(categoryID string) *CreateProductsDataBody {
	c.CategoryID = categoryID
	return c
}

func (c *CreateProductsDataBody) SetCustomData(customData interface{}) *CreateProductsDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateProductsDataBody) SetDescription(description string) *CreateProductsDataBody {
	c.Description = description
	return c
}

func (c *CreateProductsDataBody) SetFeatured(featured bool) *CreateProductsDataBody {
	c.Featured = featured
	return c
}

func (c *CreateProductsDataBody) SetImage(image string) *CreateProductsDataBody {
	c.Image = image
	return c
}

func (c *CreateProductsDataBody) SetListLabel(listLabel string) *CreateProductsDataBody {
	c.ListLabel = listLabel
	return c
}

func (c *CreateProductsDataBody) SetName(name string) *CreateProductsDataBody {
	c.Name = name
	return c
}

func (c *CreateProductsDataBody) SetOrder(order string) *CreateProductsDataBody {
	c.Order = order
	return c
}

func (c *CreateProductsDataBody) SetPrice(price string) *CreateProductsDataBody {
	c.Price = price
	return c
}

func (c *CreateProductsDataBody) SetRequiresIdentity(requiresIdentity bool) *CreateProductsDataBody {
	c.RequiresIdentity = requiresIdentity
	return c
}

func (c *CreateProductsDataBody) SetTaxRateID(taxRateID string) *CreateProductsDataBody {
	c.TaxRateID = taxRateID
	return c
}

func (c *CreateProductsDataBody) SetThumbnail(thumbnail string) *CreateProductsDataBody {
	c.Thumbnail = thumbnail
	return c
}

func (c *CreateProductsDataBody) SetURL(url string) *CreateProductsDataBody {
	c.URL = url
	return c
}

func (c *CreateProductsDataBody) SetVisible(visible bool) *CreateProductsDataBody {
	c.Visible = visible
	return c
}
