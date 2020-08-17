package slyk

import "encoding/json"

// GetCategories is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-1 TODO : Wrong
func (c Client) GetCategories(filter ...*getCategoriesFilter) (*Categories, error) {
	getBody, err := c.genericGetQuery(linkCategories, merge(filter))
	if err != nil {
		return nil, err
	}

	var categories Categories
	errUnmarshal := json.Unmarshal(getBody, &categories)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &categories, nil
}

// GetCategoryWithID is
// https://developers.slyk.io/slyk/reference/endpoints#get-categories-id
func (c Client) GetCategoryWithID(cateoryID string) (*Category, error) {
	getBody, err := c.genericGetQuery(linkCategories+"/"+cateoryID, nil)
	if err != nil {
		return nil, err
	}

	var category Category
	errUnmarshal := json.Unmarshal(getBody, &category)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &category, nil
}
