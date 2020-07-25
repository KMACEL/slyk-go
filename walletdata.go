package slyk

import (
	"strconv"
	"strings"
	"time"
)

type getWalletFilter map[string]string
type getWalletActivityFilter map[string]string
type getWalletBalanceFilter map[string]string
type getWalletMovementFilter map[string]string

type Wallets struct {
	Data  []WalletData `json:"data"`
	Total int          `json:"total"`
}

type Wallet struct {
	Data WalletData `json:"data"`
}

type WalletData struct {
	CreatedAt   time.Time   `json:"createdAt"`
	CustomData  struct{}    `json:"customData"`
	Description interface{} `json:"description"`
	ID          string      `json:"id"`
	Locked      bool        `json:"locked"`
	Metadata    struct{}    `json:"metadata"`
	Name        string      `json:"name"`
	OwnerID     string      `json:"ownerId"`
	Reference   string      `json:"reference"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type WalletActivities struct {
	Data  []WalletActivityData `json:"data"`
	Total int                  `json:"total"`
}

type WalletActivityData struct {
	Amount             string      `json:"amount"`
	AssetCode          string      `json:"assetCode"`
	Code               string      `json:"code"`
	CreatedAt          time.Time   `json:"createdAt"`
	CustomData         struct{}    `json:"customData"`
	DestinationAddress interface{} `json:"destinationAddress,omitempty"`
	DestinationWallet  struct {
		CreatedAt  time.Time `json:"createdAt"`
		CustomData struct{}  `json:"customData"`
		ID         string    `json:"id"`
		Locked     bool      `json:"locked"`
		Metadata   struct{}  `json:"metadata"`
		Name       string    `json:"name"`
		OwnerID    string    `json:"ownerId"`
		Reference  string    `json:"reference"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"destinationWallet,omitempty"`
	DestinationWalletID   string `json:"destinationWalletId"`
	DestinationWalletUser struct {
		Approved        bool      `json:"approved"`
		Blocked         bool      `json:"blocked"`
		CreatedAt       time.Time `json:"createdAt"`
		CustomData      struct{}  `json:"customData"`
		Email           string    `json:"email"`
		ID              string    `json:"id"`
		Locale          string    `json:"locale"`
		Name            string    `json:"name"`
		Phone           string    `json:"phone"`
		PrimaryWalletID string    `json:"primaryWalletId"`
		ReferralCode    string    `json:"referralCode"`
		Roles           []string  `json:"roles"`
		UpdatedAt       time.Time `json:"updatedAt"`
		Verified        bool      `json:"verified"`
	} `json:"destinationWalletUser,omitempty"`
	ExternalID     interface{} `json:"externalId"`
	ID             string      `json:"id"`
	Metadata       struct{}    `json:"metadata"`
	OriginAddress  interface{} `json:"originAddress,omitempty"`
	OriginWalletID interface{} `json:"originWalletId"`
	Status         string      `json:"status"`
	Type           string      `json:"type"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	OriginWallet   struct {
		CreatedAt  time.Time `json:"createdAt"`
		CustomData struct{}  `json:"customData"`
		ID         string    `json:"id"`
		Locked     bool      `json:"locked"`
		Metadata   struct{}  `json:"metadata"`
		Name       string    `json:"name"`
		OwnerID    string    `json:"ownerId"`
		Reference  string    `json:"reference"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"originWallet,omitempty"`
	OriginWalletUser struct {
		Approved        bool      `json:"approved"`
		Blocked         bool      `json:"blocked"`
		CreatedAt       time.Time `json:"createdAt"`
		CustomData      struct{}  `json:"customData"`
		Email           string    `json:"email"`
		ID              string    `json:"id"`
		Locale          string    `json:"locale"`
		Name            string    `json:"name"`
		Phone           string    `json:"phone"`
		PrimaryWalletID string    `json:"primaryWalletId"`
		ReferralCode    string    `json:"referralCode"`
		Roles           []string  `json:"roles"`
		UpdatedAt       time.Time `json:"updatedAt"`
		Verified        bool      `json:"verified"`
	} `json:"originWalletUser,omitempty"`
}

type WalletBalance struct {
	Data []WalletBalanceData `json:"data"`
}

type WalletBalanceData struct {
	AssetCode string `json:"assetCode"`
	Amount    string `json:"amount"`
}

type WalletMovements struct {
	Data  []WalletMovementData `json:"data"`
	Total int                  `json:"total"`
}

type WalletMovementData struct {
	Amount      string    `json:"amount"`
	AssetCode   string    `json:"assetCode"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	ID          string    `json:"id"`
	Transaction struct {
		Amount              string      `json:"amount"`
		AssetCode           string      `json:"assetCode"`
		Code                string      `json:"code"`
		CreatedAt           time.Time   `json:"createdAt"`
		CustomData          struct{}    `json:"customData"`
		Description         string      `json:"description"`
		DestinationAddress  interface{} `json:"destinationAddress"`
		DestinationWalletID string      `json:"destinationWalletId"`
		ExternalID          interface{} `json:"externalId"`
		ID                  string      `json:"id"`
		Metadata            struct {
		} `json:"metadata"`
		OriginAddress  interface{} `json:"originAddress"`
		OriginWalletID interface{} `json:"originWalletId"`
		Status         string      `json:"status"`
		Type           string      `json:"type"`
		UpdatedAt      time.Time   `json:"updatedAt"`
	} `json:"transaction"`
	TransactionID string    `json:"transactionId"`
	UpdatedAt     time.Time `json:"updatedAt"`
	WalletID      string    `json:"walletId"`
}

//=============================================================================================

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletFilter() *getWalletFilter {
	return &getWalletFilter{}
}

func (g *getWalletFilter) SetIDWithIN(id ...string) *getWalletFilter {
	(*g)["filter[id]"] = "in:" + strings.Join(id, ",")
	return g
}

func (g *getWalletFilter) SetIDWithNIN(id ...string) *getWalletFilter {
	(*g)["filter[id]"] = "nin:" + strings.Join(id, ",")
	return g
}

func (g *getWalletFilter) SetLocked(locked bool) *getWalletFilter {
	(*g)["filter[locked]"] = strconv.FormatBool(locked)
	return g
}

func (g *getWalletFilter) SetName(name string) *getWalletFilter {
	(*g)["filter[name]"] = name
	return g
}

func (g *getWalletFilter) SetOwnerID(ownerId string) *getWalletFilter {
	(*g)["filter[ownerId]"] = ownerId
	return g
}

func (g *getWalletFilter) SetReference(reference string) *getWalletFilter {
	(*g)["filter[reference]"] = reference
	return g
}

func (g *getWalletFilter) SetSortWithCreatedAt() *getWalletFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletFilter) SetSortWithCreatedAtReverse() *getWalletFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗         █████╗  ██████╗████████╗██╗██╗   ██╗██╗████████╗██╗   ██╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔════╝╚══██╔══╝██║██║   ██║██║╚══██╔══╝╚██╗ ██╔╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ███████║██║        ██║   ██║██║   ██║██║   ██║    ╚████╔╝         █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██║██║        ██║   ██║╚██╗ ██╔╝██║   ██║     ╚██╔╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║  ██║╚██████╗   ██║   ██║ ╚████╔╝ ██║   ██║      ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚═╝  ╚═══╝  ╚═╝   ╚═╝      ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletActivityFilter() *getWalletActivityFilter {
	return &getWalletActivityFilter{}
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCodeWithIN(assetCode string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletActivityFilter) SetAssetCodeWithNIN(assetCode string) *getWalletActivityFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
	return g
}

func (g *getWalletActivityFilter) SetCodeWithIN(code string) *getWalletActivityFilter {
	(*g)["filter[code]"] = "in:" + code
	return g
}

func (g *getWalletActivityFilter) SetCodeWithNIN(code string) *getWalletActivityFilter {
	(*g)["filter[code]"] = "nin:" + code
	return g
}

func (g *getWalletActivityFilter) SetStatusWithIN(status string) *getWalletActivityFilter {
	(*g)["filter[status]"] = "in:" + status
	return g
}

func (g *getWalletActivityFilter) SetStatusWithNIN(status string) *getWalletActivityFilter {
	(*g)["filter[status]"] = "nin:" + status
	return g
}

func (g *getWalletActivityFilter) SetTypeWithIN(getType string) *getWalletActivityFilter {
	(*g)["filter[type]"] = "in:" + getType
	return g
}

func (g *getWalletActivityFilter) SetTypeWithNIN(getType string) *getWalletActivityFilter {
	(*g)["filter[type]"] = "nin:" + getType
	return g
}

func (g *getWalletActivityFilter) SetSortWithAmount(amount string) *getWalletActivityFilter {
	(*g)["filter[amount]"] = "amount"
	return g
}

func (g *getWalletActivityFilter) SetSortWithAmountReverse(amount string) *getWalletActivityFilter {
	(*g)["filter[amount]"] = "-amount"
	return g
}

func (g *getWalletActivityFilter) SetSortWithCreatedAt() *getWalletActivityFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletActivityFilter) SetSortWithCreatedAtReverse() *getWalletActivityFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getWalletActivityFilter) SetAvailableIncludeWithUser() *getWalletActivityFilter {
	(*g)["include"] = "users"
	return g
}

func (g *getWalletActivityFilter) SetAvailableIncludeWithWallets() *getWalletActivityFilter {
	(*g)["include"] = "wallets"
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ██████╗  █████╗ ██╗      █████╗ ███╗   ██╗ ██████╗███████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██║     ██╔══██╗████╗  ██║██╔════╝██╔════╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██████╔╝███████║██║     ███████║██╔██╗ ██║██║     █████╗          █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██╔══██╗██╔══██║██║     ██╔══██║██║╚██╗██║██║     ██╔══╝          ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██████╔╝██║  ██║███████╗██║  ██║██║ ╚████║╚██████╗███████╗        ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletBalanceFilter() *getWalletBalanceFilter {
	return &getWalletBalanceFilter{}
}

// asset btc,eth,ltc
func (g *getWalletBalanceFilter) SetAssetCodeWithIN(assetCode string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletBalanceFilter) SetAssetCodeWithNIN(assetCode string) *getWalletBalanceFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
	return g
}

/*
 ██████╗ ███████╗████████╗        ██╗    ██╗ █████╗ ██╗     ██╗     ███████╗████████╗        ███╗   ███╗ ██████╗ ██╗   ██╗███████╗███╗   ███╗███████╗███╗   ██╗████████╗        ███████╗██╗██╗  ████████╗███████╗██████╗
██╔════╝ ██╔════╝╚══██╔══╝        ██║    ██║██╔══██╗██║     ██║     ██╔════╝╚══██╔══╝        ████╗ ████║██╔═══██╗██║   ██║██╔════╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝        ██╔════╝██║██║  ╚══██╔══╝██╔════╝██╔══██╗
██║  ███╗█████╗     ██║           ██║ █╗ ██║███████║██║     ██║     █████╗     ██║           ██╔████╔██║██║   ██║██║   ██║█████╗  ██╔████╔██║█████╗  ██╔██╗ ██║   ██║           █████╗  ██║██║     ██║   █████╗  ██████╔╝
██║   ██║██╔══╝     ██║           ██║███╗██║██╔══██║██║     ██║     ██╔══╝     ██║           ██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║           ██╔══╝  ██║██║     ██║   ██╔══╝  ██╔══██╗
╚██████╔╝███████╗   ██║           ╚███╔███╔╝██║  ██║███████╗███████╗███████╗   ██║           ██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ███████╗██║ ╚═╝ ██║███████╗██║ ╚████║   ██║           ██║     ██║███████╗██║   ███████╗██║  ██║
 ╚═════╝ ╚══════╝   ╚═╝            ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝           ╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚══════╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝           ╚═╝     ╚═╝╚══════╝╚═╝   ╚══════╝╚═╝  ╚═╝
*/

func GetWalletMovementFilter() *getWalletMovementFilter {
	return &getWalletMovementFilter{}
}

func (g *getWalletMovementFilter) SetAssetCodeWithIN(assetCode string) *getWalletMovementFilter {
	(*g)["filter[assetCode]"] = "in:" + assetCode
	return g
}

// asset btc,eth,ltc
func (g *getWalletMovementFilter) SetAssetCodeWithNIN(assetCode string) *getWalletMovementFilter {
	(*g)["filter[assetCode]"] = "nin:" + assetCode
	return g
}

func (g *getWalletMovementFilter) SetSortWithAmount(amount string) *getWalletMovementFilter {
	(*g)["filter[amount]"] = "amount"
	return g
}

func (g *getWalletMovementFilter) SetSortWithAmountReverse(amount string) *getWalletMovementFilter {
	(*g)["filter[amount]"] = "-amount"
	return g
}

func (g *getWalletMovementFilter) SetSortWithCreatedAt() *getWalletMovementFilter {
	(*g)["sort"] = "createdAt"
	return g
}

func (g *getWalletMovementFilter) SetSortWithCreatedAtReverse() *getWalletMovementFilter {
	(*g)["sort"] = "-createdAt"
	return g
}

func (g *getWalletMovementFilter) SetAvailableTransactionWithUser() *getWalletMovementFilter {
	(*g)["include"] = "transaction"
	return g
}
