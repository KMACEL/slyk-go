package slyk

import (
	"fmt"
	"strings"
)

type getaddressesFilter map[string]string

/*
 ██████╗ ███████╗████████╗         █████╗ ██████╗ ██████╗ ██████╗ ███████╗███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ███████║██║  ██║██║  ██║██████╔╝█████╗  ███████╗███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔══██║██║  ██║██║  ██║██╔══██╗██╔══╝  ╚════██║╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║  ██║██████╔╝██████╔╝██║  ██║███████╗███████║███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetAddressesFilter is
func GetAddressesFilter() *getaddressesFilter {
	return &getaddressesFilter{}
}

func (g *getaddressesFilter) SetGenericQueryParameter(key string, value interface{}) *getaddressesFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getaddressesFilter) SetAddress(address string) *getaddressesFilter {
	(*g)["address"] = address
	return g
}

// asset btc,eth,ltc
func (g *getaddressesFilter) SetAssetCode(assetCode string) *getaddressesFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getaddressesFilter) SetAssetCodeWithIN(assetCode ...string) *getaddressesFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getaddressesFilter) SetAssetCodeWithNIN(assetCode ...string) *getaddressesFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getaddressesFilter) SetID(walletID string) *getaddressesFilter {
	(*g)["filter[walletId]"] = walletID
	return g
}

func (g *getaddressesFilter) SetIDWithIN(walletID ...string) *getaddressesFilter {
	(*g)["filter[walletId]"] = "in:" + strings.Join(walletID, ",")
	return g
}

func (g *getaddressesFilter) SetIDWithNIN(walletID ...string) *getaddressesFilter {
	(*g)["filter[walletId]"] = "nin:" + strings.Join(walletID, ",")
	return g
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗         █████╗ ██████╗ ██████╗ ██████╗ ███████╗███████╗███████╗        ███████╗ ██████╗ ██████╗         ██████╗  ██████╗ ██████╗ ██╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝        ██╔════╝██╔═══██╗██╔══██╗        ██╔══██╗██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          ███████║██║  ██║██║  ██║██████╔╝█████╗  ███████╗███████╗        █████╗  ██║   ██║██████╔╝        ██████╔╝██║   ██║██║  ██║ ╚████╔╝
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██╔══██║██║  ██║██║  ██║██╔══██╗██╔══╝  ╚════██║╚════██║        ██╔══╝  ██║   ██║██╔══██╗        ██╔══██╗██║   ██║██║  ██║  ╚██╔╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ██║  ██║██████╔╝██████╔╝██║  ██║███████╗███████║███████║        ██║     ╚██████╔╝██║  ██║        ██████╔╝╚██████╔╝██████╔╝   ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝  ╚═╝╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝        ╚═╝      ╚═════╝ ╚═╝  ╚═╝        ╚═════╝  ╚═════╝ ╚═════╝    ╚═╝
*/

// CreateUserParam is
func CreateAddressForBody() *CreateAddressBody {
	return &CreateAddressBody{}
}

func (c *CreateAddressBody) SetAddress(address string) *CreateAddressBody {
	c.Address = address
	return c
}

func (c *CreateAddressBody) SetAssetCode(assetCode string) *CreateAddressBody {
	c.AssetCode = assetCode
	return c
}

func (c *CreateAddressBody) SetCustomData(customData interface{}) *CreateAddressBody {
	c.CustomData = customData
	return c
}

func (c *CreateAddressBody) SetProvider(provider string) *CreateAddressBody {
	c.Provider = provider
	return c
}

func (c *CreateAddressBody) SetWalletID(walletID string) *CreateAddressBody {
	c.WalletID = walletID
	return c
}
