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

	returnValue, err := getClient().UpdateCategory("27d3d24a-257b-4a88-a8eb-bb8b6121f317", &UpdateCategoryDataBody{
		CustomData: map[string]interface{}{"Acel": "test"},
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateCategory(t *testing.T) {
	tst := "TestCreateCategory"

	returnValue, err := getClient().CreateCategory(CreateCategoryDataForBody().SetCustomData(map[string]interface{}{"mert": "category"}).SetTitle("Test1"))

	ReturnAndError(t, tst, returnValue, err)
}

func TestCategoryReorder(t *testing.T) {
	tst := "TestCategoryReorder"

	err := getClient().CategoryReorder("27d3d24a-257b-4a88-a8eb-bb8b6121f317", &CategoryReorderDataBody{})

	OnlyError(t, tst, err)
}

func TestDeleteCategory(t *testing.T) {
	tst := "TestDeleteCategory"

	err := getClient().DeleteCategory("ec3cf024-0f6d-4252-bd0b-86795d871482")

	OnlyError(t, tst, err)
}
