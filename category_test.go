package slyk

import "testing"

func TestGetCategories(t *testing.T) {
	tst := "TestGetCategories"

	returnValue, err := getClient().GetCategories()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetCategoryWithID(t *testing.T) {
	tst := "TestGetCategoryWithID"

	returnValue, err := getClient().GetCategoryWithID("27d3d24a-257b-4a88-a8eb-bb8b6121f317")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateCategoryWithID(t *testing.T) {
	tst := "TestUpdateCategoryWithID"

	returnValue, err := getClient().UpdateCategoryWithID("27d3d24a-257b-4a88-a8eb-bb8b6121f317", &UpdateCategoryDataBody{
		CustomData: map[string]interface{}{"Acel": "test"},
	})

	ReturnAndError(t, tst, returnValue, err)
}
