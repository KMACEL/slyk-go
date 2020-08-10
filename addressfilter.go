package slyk

import (
	"fmt"
	"strings"
)

type getaddressFilter map[string]string

/*
 ██████╗ ███████╗████████╗         █████╗ ██████╗ ██████╗ ██████╗ ███████╗███████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ███████║██║  ██║██║  ██║██████╔╝█████╗  ███████╗███████╗        █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██╔══██║██║  ██║██║  ██║██╔══██╗██╔══╝  ╚════██║╚════██║        ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ██║  ██║██████╔╝██████╔╝██║  ██║███████╗███████║███████║        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═╝  ╚═╝╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// GetAddressFilter is
func GetAddressFilter() *getaddressFilter {
	return &getaddressFilter{}
}

func (g *getaddressFilter) SetGenericQueryParameter(key string, value interface{}) *getaddressFilter {
	(*g)[key] = fmt.Sprintf("%v", value)
	return g
}

func (g *getaddressFilter) SetAddress(address string) *getaddressFilter {
	(*g)["address"] = address
	return g
}

// asset btc,eth,ltc
func (g *getaddressFilter) SetAssetCode(assetCode string) *getaddressFilter {
	(*g)["filter[assetCode]"] = assetCode
	return g
}

// asset btc,eth,ltc
func (g *getaddressFilter) SetAssetCodeWithIN(assetCode ...string) *getaddressFilter {
	(*g)["filter[assetCode]"] = "in:" + strings.Join(assetCode, ",")
	return g
}

// asset btc,eth,ltc
func (g *getaddressFilter) SetAssetCodeWithNIN(assetCode ...string) *getaddressFilter {
	(*g)["filter[assetCode]"] = "nin:" + strings.Join(assetCode, ",")
	return g
}

func (g *getaddressFilter) SetID(walletID string) *getaddressFilter {
	(*g)["filter[walletId]"] = walletID
	return g
}

func (g *getaddressFilter) SetIDWithIN(walletID ...string) *getaddressFilter {
	(*g)["filter[walletId]"] = "in:" + strings.Join(walletID, ",")
	return g
}

func (g *getaddressFilter) SetIDWithNIN(walletID ...string) *getaddressFilter {
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
