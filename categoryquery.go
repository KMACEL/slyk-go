package slyk

import "encoding/json"

// GetCategories is
// https://developers.slyk.io/slyk/reference/endpoints#get-invites-1 TODO : Wrong
func (c Client) GetCategories(filter ...*getCategoriesFilter) (*Categories, error) {
	getBody, err := c.GenericGetQuery(linkCategories, merge(filter))
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
	getBody, err := c.GenericGetQuery(linkCategories+"/"+cateoryID, nil)
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

// UpdateCategoryWithID is
// https://developers.slyk.io/slyk/reference/endpoints#patch-categories-id is
func (c Client) UpdateCategory(cateoryID string, updateCategoryDataBody *UpdateCategoryDataBody) (*Category, error) {
	getBody, err := c.GenericPatchQuery(linkCategories+"/"+cateoryID, updateCategoryDataBody)
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

// CreateCategory is
// https://developers.slyk.io/slyk/reference/endpoints#post-categories
func (c Client) CreateCategory(createCategoryDataBody *CreateCategoryDataBody) (*Category, error) {
	getBody, err := c.GenericPostQuery(linkCategories, createCategoryDataBody)
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

// CategoryReorder is
// https://developers.slyk.io/slyk/reference/endpoints#post-categories-id-reorder
func (c Client) CategoryReorder(cateoryID string, categoryReorderDataBody *CategoryReorderDataBody) error {
	_, err := c.GenericPostQuery(linkCategories+"/"+cateoryID+"/reorder", categoryReorderDataBody)
	return err
}

// https://developers.slyk.io/slyk/reference/endpoints#delete-categories-id
func (c Client) DeleteCategory(cateoryID string) error {
	return c.GenericDeleteQuery(linkCategories+"/"+cateoryID, nil)
}
