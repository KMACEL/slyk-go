package slyk

import (
	"encoding/json"
)

// GetAssets is
// https://developers.slyk.io/slyk/reference/endpoints#get-assets
func (c Client) GetAssets(filter ...*getassetsFilter) (*Assests, error) {
	getBody, err := c.GenericGetQuery(linkAssets, merge(filter))
	if err != nil {
		return nil, err
	}

	var assests Assests
	errUnmarshal := json.Unmarshal(getBody, &assests)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &assests, nil
}

// GetAssetsWithCode is
// https://developers.slyk.io/slyk/reference/endpoints#get-assets-code
func (c Client) GetAssetsWithCode(assetCode string) (*Asset, error) {
	getBody, err := c.GenericGetQuery(linkAssets+"/"+assetCode, nil)
	if err != nil {
		return nil, err
	}

	var asset Asset
	errUnmarshal := json.Unmarshal(getBody, &asset)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &asset, nil
}

// UpdateAssetsWithCode
// https://developers.slyk.io/slyk/reference/endpoints#patch-assets-code
func (c Client) UpdateAssets(assetCode string, updateAssetDataBody *UpdateAssetDataBody) (*Asset, error) {
	getBody, err := c.GenericPatchQuery(linkAssets+"/"+assetCode, updateAssetDataBody)
	if err != nil {
		return nil, err
	}

	var asset Asset
	errUnmarshal := json.Unmarshal(getBody, &asset)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &asset, nil
}

// CreateAsset
// https://developers.slyk.io/slyk/reference/endpoints#post-assets
func (c Client) CreateAsset(createAssetDataBody *CreateAssetDataBody) (*Asset, error) {
	getBody, err := c.GenericPostQuery(linkAssets, createAssetDataBody)
	if err != nil {
		return nil, err
	}

	var asset Asset
	errUnmarshal := json.Unmarshal(getBody, &asset)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &asset, nil
}
