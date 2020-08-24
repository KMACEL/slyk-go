package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getQuestionsFilter map[string]string
type getQuestionsTypesFilter map[string]string

/*
 ██████╗ ███████╗████████╗         ██████╗ ██╗   ██╗███████╗███████╗████████╗██╗ ██████╗ ███╗   ██╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔═══██╗██║   ██║██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║   ██║██║   ██║█████╗  ███████╗   ██║   ██║██║   ██║██╔██╗ ██║███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║▄▄ ██║██║   ██║██╔══╝  ╚════██║   ██║   ██║██║   ██║██║╚██╗██║╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚██████╔╝╚██████╔╝███████╗███████║   ██║   ██║╚██████╔╝██║ ╚████║███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══▀▀═╝  ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetQuestionsFilter is
func GetQuestionsFilter() *getQuestionsFilter {
	return &getQuestionsFilter{}
}

func (g *getQuestionsFilter) SetGenericQueryParameter(key string, value interface{}) *getQuestionsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getQuestionsFilter) SetID(id string) *getQuestionsFilter {
	(*g)["id"] = id
	return g
}

func (g *getQuestionsFilter) SetIDWithIN(id ...string) *getQuestionsFilter {
	(*g)["id"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getQuestionsFilter) SetIDWithNIN(id ...string) *getQuestionsFilter {
	(*g)["id"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getQuestionsFilter) SetProductTypeCode(productTypeCode string) *getQuestionsFilter {
	(*g)["productTypeCode"] = productTypeCode
	return g
}

func (g *getQuestionsFilter) SetTitle(title string) *getQuestionsFilter {
	(*g)["title"] = title
	return g
}

func (g *getQuestionsFilter) SetTypeCode(typeCode string) *getQuestionsFilter {
	(*g)["typeCode"] = typeCode
	return g
}

func (g *getQuestionsFilter) SetTypeCodeWithIN(typeCode ...string) *getQuestionsFilter {
	(*g)["typeCode"] = "in:" + strings.Join(typeCode, ",")
	return g
}

func (g *getQuestionsFilter) SetTypeCodeWithNIN(typeCode ...string) *getQuestionsFilter {
	(*g)["typeCode"] = "nin:" + strings.Join(typeCode, ",")
	return g
}

func (g *getQuestionsFilter) SetSortWithCreatedAt() *getQuestionsFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getQuestionsFilter) SetSortWithCreatedAtReverse() *getQuestionsFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getQuestionsFilter) SetSortWithTypeCode() *getQuestionsFilter {
	(*g)["sort"] = "typeCode"
	return g
}

func (g *getQuestionsFilter) SetSortWithTypeCodeReverse() *getQuestionsFilter {
	(*g)["sort"] = "-typeCode"
	return g
}

func (g *getQuestionsFilter) SetSortWithTitle() *getQuestionsFilter {
	(*g)["sort"] = "title"
	return g
}

func (g *getQuestionsFilter) SetSortWithTitleReverse() *getQuestionsFilter {
	(*g)["sort"] = "-title"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getQuestionsFilter) SetPageSize(size int) *getQuestionsFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getQuestionsFilter) SetPageNumber(number int) *getQuestionsFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
 ██████╗ ███████╗████████╗         ██████╗ ██╗   ██╗███████╗███████╗████████╗██╗ ██████╗ ███╗   ██╗███████╗        ████████╗██╗   ██╗██████╗ ███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔═══██╗██║   ██║██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝        ╚══██╔══╝╚██╗ ██╔╝██╔══██╗██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║   ██║██║   ██║█████╗  ███████╗   ██║   ██║██║   ██║██╔██╗ ██║███████╗           ██║    ╚████╔╝ ██████╔╝█████╗  ███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║▄▄ ██║██║   ██║██╔══╝  ╚════██║   ██║   ██║██║   ██║██║╚██╗██║╚════██║           ██║     ╚██╔╝  ██╔═══╝ ██╔══╝  ╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚██████╔╝╚██████╔╝███████╗███████║   ██║   ██║╚██████╔╝██║ ╚████║███████║           ██║      ██║   ██║     ███████╗███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══▀▀═╝  ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝           ╚═╝      ╚═╝   ╚═╝     ╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetQuestionsTypesFilter is
func GetQuestionsTypesFilter() *getQuestionsTypesFilter {
	return &getQuestionsTypesFilter{}
}

func (g *getQuestionsTypesFilter) SetGenericQueryParameter(key string, value interface{}) *getQuestionsTypesFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getQuestionsTypesFilter) SetCode(code string) *getQuestionsTypesFilter {
	(*g)["code"] = code
	return g
}

func (g *getQuestionsTypesFilter) SetCodeWithIN(code ...string) *getQuestionsTypesFilter {
	(*g)["code"] = "in:" + strings.Join(code, ",")
	return g
}

func (g *getQuestionsTypesFilter) SetCodeWithNIN(code ...string) *getQuestionsTypesFilter {
	(*g)["code"] = "nin:" + strings.Join(code, ",")
	return g
}

func (g *getQuestionsTypesFilter) SetSortWithCreatedAt() *getQuestionsTypesFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getQuestionsTypesFilter) SetSortWithCreatedAtReverse() *getQuestionsTypesFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getQuestionsTypesFilter) SetSortWithCode() *getQuestionsTypesFilter {
	(*g)["sort"] = "code"
	return g
}

func (g *getQuestionsTypesFilter) SetSortWithCodeReverse() *getQuestionsTypesFilter {
	(*g)["sort"] = "-code"
	return g
}

// Defines the number of results per page. Default = 30.
func (g *getQuestionsTypesFilter) SetPageSize(size int) *getQuestionsTypesFilter {
	(*g)["page[size]"] = strconv.Itoa(size)
	return g
}

// Defines the number of the page to retrieve. Default = 1
func (g *getQuestionsTypesFilter) SetPageNumber(number int) *getQuestionsTypesFilter {
	(*g)["page[number]"] = strconv.Itoa(number)
	return g
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗         ██████╗ ██╗   ██╗███████╗███████╗████████╗██╗ ██████╗ ███╗   ██╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝        ██╔═══██╗██║   ██║██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗          ██║   ██║██║   ██║█████╗  ███████╗   ██║   ██║██║   ██║██╔██╗ ██║        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝          ██║▄▄ ██║██║   ██║██╔══╝  ╚════██║   ██║   ██║██║   ██║██║╚██╗██║        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗        ╚██████╔╝╚██████╔╝███████╗███████║   ██║   ██║╚██████╔╝██║ ╚████║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚══▀▀═╝  ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// UpdateQuestionDataForBody is
func UpdateQuestionDataForBody() *UpdateQuestionDataBody {
	return &UpdateQuestionDataBody{}
}

func (u *UpdateQuestionDataBody) SetConfigurations(configurations interface{}) *UpdateQuestionDataBody {
	u.Configurations = configurations
	return u
}

func (u *UpdateQuestionDataBody) SetCustomData(customData interface{}) *UpdateQuestionDataBody {
	u.CustomData = customData
	return u
}

func (u *UpdateQuestionDataBody) SetDescription(description string) *UpdateQuestionDataBody {
	u.Description = description
	return u
}

func (u *UpdateQuestionDataBody) SetProductTypeCode(productTypeCode string) *UpdateQuestionDataBody {
	u.ProductTypeCode = productTypeCode
	return u
}

func (u *UpdateQuestionDataBody) SetRequired(required bool) *UpdateQuestionDataBody {
	u.Required = required
	return u
}

func (u *UpdateQuestionDataBody) SetTitle(title string) *UpdateQuestionDataBody {
	u.Title = title
	return u
}

func (u *UpdateQuestionDataBody) SetTypeCode(typeCode string) *UpdateQuestionDataBody {
	u.TypeCode = typeCode
	return u
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗         ██████╗ ██╗   ██╗███████╗███████╗████████╗██╗ ██████╗ ███╗   ██╗        ██████╗  █████╗ ████████╗ █████╗         ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔═══██╗██║   ██║██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║        ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ██║   ██║██║   ██║█████╗  ███████╗   ██║   ██║██║   ██║██╔██╗ ██║        ██║  ██║███████║   ██║   ███████║        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██║▄▄ ██║██║   ██║██╔══╝  ╚════██║   ██║   ██║██║   ██║██║╚██╗██║        ██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ╚██████╔╝╚██████╔╝███████╗███████║   ██║   ██║╚██████╔╝██║ ╚████║        ██████╔╝██║  ██║   ██║   ██║  ██║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝         ╚══▀▀═╝  ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CreateQuestionDataForBody is
func CreateQuestionDataForBody() *CreateQuestionDataBody {
	return &CreateQuestionDataBody{}
}

func (c *CreateQuestionDataBody) SetConfigurations(configurations interface{}) *CreateQuestionDataBody {
	c.Configurations = configurations
	return c
}

func (c *CreateQuestionDataBody) SetCustomData(customData interface{}) *CreateQuestionDataBody {
	c.CustomData = customData
	return c
}

func (c *CreateQuestionDataBody) SetDescription(description string) *CreateQuestionDataBody {
	c.Description = description
	return c
}

func (c *CreateQuestionDataBody) SetProductTypeCode(productTypeCode string) *CreateQuestionDataBody {
	c.ProductTypeCode = productTypeCode
	return c
}

func (c *CreateQuestionDataBody) SetRequired(required bool) *CreateQuestionDataBody {
	c.Required = required
	return c
}

func (c *CreateQuestionDataBody) SetTitle(title string) *CreateQuestionDataBody {
	c.Title = title
	return c
}

func (c *CreateQuestionDataBody) SetTypeCode(typeCode string) *CreateQuestionDataBody {
	c.TypeCode = typeCode
	return c
}
