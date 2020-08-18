package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getPaymentMethodsFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██████╗  █████╗ ██╗   ██╗███╗   ███╗███████╗███╗   ██╗████████╗        ███╗   ███╗███████╗██████╗ ██╗  ██╗ ██████╗ ██████╗         ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗╚██╗ ██╔╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ████╗ ████║██╔════╝██╔══██╗██║  ██║██╔═══██╗██╔══██╗        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██████╔╝███████║ ╚████╔╝ ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           ██╔████╔██║█████╗  ██║  ██║███████║██║   ██║██║  ██║        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔═══╝ ██╔══██║  ╚██╔╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██║╚██╔╝██║██╔══╝  ██║  ██║██╔══██║██║   ██║██║  ██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║     ██║  ██║   ██║   ██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ██║ ╚═╝ ██║███████╗██████╔╝██║  ██║╚██████╔╝██████╔╝        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═════╝         ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetPaymentMedhodFilter is
func GetPaymentMedhodsFilter() *getPaymentMethodsFilter {
	return &getPaymentMethodsFilter{}
}

func (g *getPaymentMethodsFilter) SetGenericQueryParameter(key string, value interface{}) *getPaymentMethodsFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// Format : 2019-07-21
func (g *getPaymentMethodsFilter) SetCreatedAtWithGTE(date string) *getPaymentMethodsFilter {
	(*g)["filter[createdAt]"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getPaymentMethodsFilter) SetCreatedAtWithLTE(date string) *getPaymentMethodsFilter {
	(*g)["filter[createdAt]"] = "lte:" + date
	return g
}

func (g *getPaymentMethodsFilter) SetEnabled(enabled bool) *getPaymentMethodsFilter {
	(*g)["filter[enabled]"] = strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodsFilter) SetEnabledWithIN(enabled bool) *getPaymentMethodsFilter {
	(*g)["filter[enabled]"] = "in:" + strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodsFilter) SetEnabledWithNIN(enabled bool) *getPaymentMethodsFilter {
	(*g)["filter[enabled]"] = "nin:" + strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodsFilter) SetSlug(slug string) *getPaymentMethodsFilter {
	(*g)["filter[slug]"] = slug
	return g
}

func (g *getPaymentMethodsFilter) SetSlugWithIN(slug ...string) *getPaymentMethodsFilter {
	(*g)["filter[slug]"] = "in:" + strings.Join(slug, ",")
	return g
}

func (g *getPaymentMethodsFilter) SetSlugWithNIN(slug ...string) *getPaymentMethodsFilter {
	(*g)["filter[slug]"] = "nin:" + strings.Join(slug, ",")
	return g
}
