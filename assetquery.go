package slyk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// GetAssets is
// https://developers.slyk.io/slyk/reference/endpoints#get-assets
func (c Client) GetAssets(filter ...*getassetFilter) (*Assests, error) {

	clientReq := resty.New().R()
	if filter != nil {
		clientReq.SetQueryParams(merge(filter))
	}

	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		Get(linkAssets)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var assests Assests
	errUnmarshal := json.Unmarshal(resp.Body(), &assests)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &assests, nil
}

// GetAssetsWithCode is
// https://developers.slyk.io/slyk/reference/endpoints#get-assets-code
func (c Client) GetAssetsWithCode(assetCode string) (*Asset, error) {

	resp, err := resty.New().R().
		SetHeader(headerApiKey, c.apiKey).
		Get(linkAssets + "/" + assetCode)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var asset Asset
	errUnmarshal := json.Unmarshal(resp.Body(), &asset)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &asset, nil
}

// UpdateAssetsWithCode
// https://developers.slyk.io/slyk/reference/endpoints#patch-assets-code
func (c Client) UpdateAssetsWithCode(assetCode string, updateAssetDataBody *UpdateAssetDataBody) (*Asset, error) {
	clientReq := resty.New().R()
	resp, err := clientReq.
		SetHeader(headerApiKey, c.apiKey).
		SetBody(updateAssetDataBody).
		Patch(linkAssets + "/" + assetCode)

	fmt.Println(clientReq)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Status Code : %d", resp.StatusCode())
	}

	var asset Asset
	errUnmarshal := json.Unmarshal(resp.Body(), &asset)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &asset, nil
}
