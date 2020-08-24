package slyk

import "time"

type Products struct {
	Data  []ProductData `json:"data"`
	Total int           `json:"total"`
}

type Product struct {
	Data ProductData `json:"data"`
}

type ProductData struct {
	AllowChoosingQuantity bool      `json:"allowChoosingQuantity"`
	AssetCode             string    `json:"assetCode"`
	Available             bool      `json:"available"`
	Bonus                 string    `json:"bonus"`
	ButtonLabel           string    `json:"buttonLabel"`
	CategoryID            string    `json:"categoryId"`
	CreatedAt             time.Time `json:"createdAt"`
	CustomData            struct {
	} `json:"customData"`
	Description string `json:"description"`
	Featured    bool   `json:"featured"`
	ID          string `json:"id"`
	ImageURL    string `json:"imageUrl"`
	ListLabel   string `json:"listLabel"`
	Metadata    struct {
	} `json:"metadata"`
	Name             string    `json:"name"`
	Price            string    `json:"price"`
	PriceWithTax     string    `json:"priceWithTax"`
	RequiresIdentity bool      `json:"requiresIdentity"`
	TaxAmount        string    `json:"taxAmount"`
	TaxRateID        string    `json:"taxRateId"`
	ThumbnailURL     string    `json:"thumbnailUrl"`
	UpdatedAt        time.Time `json:"updatedAt"`
	URL              string    `json:"url"`
	Visible          bool      `json:"visible"`
}

type CreateProductDataBody struct {
	AllowChoosingQuantity bool        `json:"allowChoosingQuantity,omitempty"`
	AssetCode             string      `json:"assetCode,omitempty"`
	Available             bool        `json:"available,omitempty"`
	Bonus                 string      `json:"bonus,omitempty"`
	ButtonLabel           string      `json:"buttonLabel,omitempty"`
	CategoryID            string      `json:"categoryId"`
	CustomData            interface{} `json:"customData,omitempty"`
	Description           string      `json:"description,omitempty"`
	Featured              bool        `json:"featured,omitempty"`
	Image                 string      `json:"image,omitempty"`
	ListLabel             string      `json:"listLabel,omitempty"`
	Name                  string      `json:"name"`
	Order                 string      `json:"order,omitempty"`
	Price                 string      `json:"price"`
	RequiresIdentity      bool        `json:"requiresIdentity,omitempty"`
	TaxRateID             string      `json:"taxRateId,omitempty"`
	Thumbnail             string      `json:"thumbnail,omitempty"`
	URL                   string      `json:"url,omitempty"`
	Visible               bool        `json:"visible,omitempty"`
}

type UpdateProductDataBody struct {
	AllowChoosingQuantity bool        `json:"allowChoosingQuantity,omitempty"`
	AssetCode             string      `json:"assetCode,omitempty"`
	Available             bool        `json:"available,omitempty"`
	Bonus                 string      `json:"bonus,omitempty"`
	ButtonLabel           string      `json:"buttonLabel,omitempty"`
	CategoryID            string      `json:"categoryId,omitempty"`
	CustomData            interface{} `json:"customData,omitempty"`
	Description           string      `json:"description,omitempty"`
	Featured              bool        `json:"featured,omitempty"`
	Image                 string      `json:"image,omitempty"`
	ListLabel             string      `json:"listLabel,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Order                 string      `json:"order,omitempty"`
	Price                 string      `json:"price,omitempty"`
	RequiresIdentity      bool        `json:"requiresIdentity,omitempty"`
	TaxRateID             string      `json:"taxRateId,omitempty"`
	Thumbnail             string      `json:"thumbnail,omitempty"`
	URL                   string      `json:"url,omitempty"`
	Visible               bool        `json:"visible,omitempty"`
}

type AddProductQuestionDataBody struct {
	QuestionID string `json:"questionId"`
}

type AddProductQuestionResponseBody struct {
	Data struct {
		Order      string `json:"order"`
		ProductID  string `json:"productId"`
		QuestionID string `json:"questionId"`
	} `json:"data"`
}

type ProductReorderDataBody struct {
	SubsequentID string `json:"subsequentId"`
}

type ProductQuestionReorderDataBody struct {
	SubsequentID string `json:"subsequentId"`
}
