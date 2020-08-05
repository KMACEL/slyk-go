package slyk

import (
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

// Format : 2019-07-21
func (g *getPaymentMethodFilter) SetCreatedAtWithGTE(date string) *getPaymentMethodFilter {
	(*g)["createdAt"] = "gte:" + date
	return g
}

// Format : 2019-07-21
func (g *getPaymentMethodFilter) SetCreatedAtWithLTE(date string) *getPaymentMethodFilter {
	(*g)["createdAt"] = "lte:" + date
	return g
}

func (g *getPaymentMethodFilter) SetEnabledWithIN(enabled bool) *getPaymentMethodFilter {
	(*g)["filter[enabled]"] = strconv.FormatBool(enabled)
	return g
}

func (g *getPaymentMethodFilter) SetEnabledWithNIN(enabled bool) *getPaymentMethodFilter {
	(*g)["filter[enabled]"] = strconv.FormatBool(enabled)
	return g
}

func (g *getWalletFilter) SetSlugWithIN(slug ...string) *getWalletFilter {
	(*g)["filter[slug]"] = "in:" + strings.Join(slug, ",")
	return g
}

func (g *getWalletFilter) SetSlugWithNIN(slug ...string) *getWalletFilter {
	(*g)["filter[slug]"] = "nin:" + strings.Join(slug, ",")
	return g
}
