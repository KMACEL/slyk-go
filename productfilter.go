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
func CreateProductDataForBody() *CreateProductDataBody {
	return &CreateProductDataBody{}
}

func (c *CreateProductDataBody) SetAllowChoosingQuantity(allowChoosingQuantity bool) *CreateProductDataBody {
	c.AllowChoosingQuantity = allowChoosingQuantity
	return c
}

func (c *CreateProductDataBody) SetAssetCode(assetCode string) *CreateProductDataBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateProductDataBody) SetAvailable(available bool) *CreateProductDataBody {
	c.Available = available
	return c
}

func (c *CreateProductDataBody) SetBonus(bonus string) *CreateProductDataBody {
	c.Bonus = bonus
	return c
}

func (c *CreateProductDataBody) SetButtonLabel(buttonLabel string) *CreateProductDataBody {
	c.ButtonLabel = buttonLabel
	return c
}

func (c *CreateProductDataBody) SetCategoryID(categoryID string) *CreateProductDataBody {
	c.CategoryID = categoryID
	return c
}

func (c *CreateProductDataBody) SetCustomData(customData interface{}) *CreateProductDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateProductDataBody) SetDescription(description string) *CreateProductDataBody {
	c.Description = description
	return c
}

func (c *CreateProductDataBody) SetFeatured(featured bool) *CreateProductDataBody {
	c.Featured = featured
	return c
}

func (c *CreateProductDataBody) SetImage(image string) *CreateProductDataBody {
	c.Image = image
	return c
}

func (c *CreateProductDataBody) SetListLabel(listLabel string) *CreateProductDataBody {
	c.ListLabel = listLabel
	return c
}

func (c *CreateProductDataBody) SetName(name string) *CreateProductDataBody {
	c.Name = name
	return c
}

func (c *CreateProductDataBody) SetOrder(order string) *CreateProductDataBody {
	c.Order = order
	return c
}

func (c *CreateProductDataBody) SetPrice(price string) *CreateProductDataBody {
	c.Price = price
	return c
}

func (c *CreateProductDataBody) SetRequiresIdentity(requiresIdentity bool) *CreateProductDataBody {
	c.RequiresIdentity = requiresIdentity
	return c
}

func (c *CreateProductDataBody) SetTaxRateID(taxRateID string) *CreateProductDataBody {
	c.TaxRateID = taxRateID
	return c
}

func (c *CreateProductDataBody) SetThumbnail(thumbnail string) *CreateProductDataBody {
	c.Thumbnail = thumbnail
	return c
}

func (c *CreateProductDataBody) SetURL(url string) *CreateProductDataBody {
	c.URL = url
	return c
}

func (c *CreateProductDataBody) SetVisible(visible bool) *CreateProductDataBody {
	c.Visible = visible
	return c
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗        ██████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗ ██████╗████████╗███████╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██║   ██║██╔════╝╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗          ██████╔╝██████╔╝██║   ██║██║  ██║██║   ██║██║        ██║   ███████╗        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝          ██╔═══╝ ██╔══██╗██║   ██║██║  ██║██║   ██║██║        ██║   ╚════██║        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗        ██║     ██║  ██║╚██████╔╝██████╔╝╚██████╔╝╚██████╗   ██║   ███████║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝   ╚═╝   ╚══════╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// UpdateProductsDataForBody is
func UpdateProductDataForBody() *UpdateProductDataBody {
	return &UpdateProductDataBody{}
}

func (u *UpdateProductDataBody) SetAllowChoosingQuantity(allowChoosingQuantity bool) *UpdateProductDataBody {
	u.AllowChoosingQuantity = allowChoosingQuantity
	return u
}

func (u *UpdateProductDataBody) SetAssetCode(assetCode string) *UpdateProductDataBody {
	u.AssetCode = assetCode
	return u
}

func (u *UpdateProductDataBody) SetAvailable(available bool) *UpdateProductDataBody {
	u.Available = available
	return u
}

func (u *UpdateProductDataBody) SetBonus(bonus string) *UpdateProductDataBody {
	u.Bonus = bonus
	return u
}

func (u *UpdateProductDataBody) SetButtonLabel(buttonLabel string) *UpdateProductDataBody {
	u.ButtonLabel = buttonLabel
	return u
}

func (u *UpdateProductDataBody) SetCategoryID(categoryID string) *UpdateProductDataBody {
	u.CategoryID = categoryID
	return u
}

func (u *UpdateProductDataBody) SetCustomData(customData interface{}) *UpdateProductDataBody {
	u.CustomData = customData
	return u
}

func (u *UpdateProductDataBody) SetDescription(description string) *UpdateProductDataBody {
	u.Description = description
	return u
}

func (u *UpdateProductDataBody) SetFeatured(featured bool) *UpdateProductDataBody {
	u.Featured = featured
	return u
}

func (u *UpdateProductDataBody) SetImage(image string) *UpdateProductDataBody {
	u.Image = image
	return u
}

func (u *UpdateProductDataBody) SetListLabel(listLabel string) *UpdateProductDataBody {
	u.ListLabel = listLabel
	return u
}

func (u *UpdateProductDataBody) SetName(name string) *UpdateProductDataBody {
	u.Name = name
	return u
}

func (u *UpdateProductDataBody) SetOrder(order string) *UpdateProductDataBody {
	u.Order = order
	return u
}

func (u *UpdateProductDataBody) SetPrice(price string) *UpdateProductDataBody {
	u.Price = price
	return u
}

func (u *UpdateProductDataBody) SetRequiresIdentity(requiresIdentity bool) *UpdateProductDataBody {
	u.RequiresIdentity = requiresIdentity
	return u
}

func (u *UpdateProductDataBody) SetTaxRateID(taxRateID string) *UpdateProductDataBody {
	u.TaxRateID = taxRateID
	return u
}

func (u *UpdateProductDataBody) SetThumbnail(thumbnail string) *UpdateProductDataBody {
	u.Thumbnail = thumbnail
	return u
}

func (u *UpdateProductDataBody) SetURL(url string) *UpdateProductDataBody {
	u.URL = url
	return u
}

func (u *UpdateProductDataBody) SetVisible(visible bool) *UpdateProductDataBody {
	u.Visible = visible
	return u
}

/*
 █████╗ ██████╗ ██████╗         ██████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗ ██████╗████████╗         ██████╗ ██╗   ██╗███████╗███████╗████████╗██╗ ██████╗ ███╗   ██╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔══██╗██╔══██╗        ██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██║   ██║██╔════╝╚══██╔══╝        ██╔═══██╗██║   ██║██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
███████║██║  ██║██║  ██║        ██████╔╝██████╔╝██║   ██║██║  ██║██║   ██║██║        ██║           ██║   ██║██║   ██║█████╗  ███████╗   ██║   ██║██║   ██║██╔██╗ ██║        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██╔══██║██║  ██║██║  ██║        ██╔═══╝ ██╔══██╗██║   ██║██║  ██║██║   ██║██║        ██║           ██║▄▄ ██║██║   ██║██╔══╝  ╚════██║   ██║   ██║██║   ██║██║╚██╗██║        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
██║  ██║██████╔╝██████╔╝        ██║     ██║  ██║╚██████╔╝██████╔╝╚██████╔╝╚██████╗   ██║           ╚██████╔╝╚██████╔╝███████╗███████║   ██║   ██║╚██████╔╝██║ ╚████║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
╚═╝  ╚═╝╚═════╝ ╚═════╝         ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝   ╚═╝            ╚══▀▀═╝  ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// AddProductQuestionDataForBody is
func AddProductQuestionDataForBody() *AddProductQuestionDataBody {
	return &AddProductQuestionDataBody{}
}

func (u *AddProductQuestionDataBody) SetQuestionID(questionID string) *AddProductQuestionDataBody {
	u.QuestionID = questionID
	return u
}
