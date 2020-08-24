package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getQuestionsFilter map[string]string

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
