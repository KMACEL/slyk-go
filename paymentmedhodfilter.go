package slyk

import (
	"fmt"
	"strconv"
	"strings"
)

type getPaymentMethodFilter map[string]string

/*
 ██████╗ ███████╗████████╗        ██████╗  █████╗ ██╗   ██╗███╗   ███╗███████╗███╗   ██╗████████╗        ███╗   ███╗███████╗██████╗ ██╗  ██╗ ██████╗ ██████╗         ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗╚██╗ ██╔╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ████╗ ████║██╔════╝██╔══██╗██║  ██║██╔═══██╗██╔══██╗        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██████╔╝███████║ ╚████╔╝ ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           ██╔████╔██║█████╗  ██║  ██║███████║██║   ██║██║  ██║        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔═══╝ ██╔══██║  ╚██╔╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██║╚██╔╝██║██╔══╝  ██║  ██║██╔══██║██║   ██║██║  ██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║     ██║  ██║   ██║   ██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ██║ ╚═╝ ██║███████╗██████╔╝██║  ██║╚██████╔╝██████╔╝        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═════╝         ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetPaymentMedhodFilter is
func GetPaymentMedhodFilter() *getPaymentMethodFilter {
	return &getPaymentMethodFilter{}
}

func (g *getPaymentMethodFilter) SetGenericQueryParameter(key string, value interface{}) *getPaymentMethodFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

// Format : 2019-07-21
func (g *getPaymentMethodFilter) SetCreatedAtWithGTE(date string) *getPaymentMethodFilter {
	(*g)["filter[createdAt]"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getPaymentMethodFilter) SetCreatedAtWithLTE(date string) *getPaymentMethodFilter {
	(*g)["filter[createdAt]"] = "lte:" + date
	return g
}

func (g *getPaymentMethodFilter) SetEnabled(enabled bool) *getPaymentMethodFilter {
	(*g)["filter[enabled]"] = strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodFilter) SetEnabledWithIN(enabled bool) *getPaymentMethodFilter {
	(*g)["filter[enabled]"] = "in:" + strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodFilter) SetEnabledWithNIN(enabled bool) *getPaymentMethodFilter {
	(*g)["filter[enabled]"] = "nin:" + strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodFilter) SetSlug(slug string) *getPaymentMethodFilter {
	(*g)["filter[slug]"] = slug
	return g
}

func (g *getPaymentMethodFilter) SetSlugWithIN(slug ...string) *getPaymentMethodFilter {
	(*g)["filter[slug]"] = "in:" + strings.Join(slug, ",")
	return g
}

func (g *getPaymentMethodFilter) SetSlugWithNIN(slug ...string) *getPaymentMethodFilter {
	(*g)["filter[slug]"] = "nin:" + strings.Join(slug, ",")
	return g
}
