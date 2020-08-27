package slyk

import "testing"

func TestGetProducts(t *testing.T) {
	tst := "TestGetProducts"

	returnValue, err := getClient().GetProducts()

	ReturnAndError(t, tst, returnValue, err)
}

func TestGetProductWithID(t *testing.T) {
	tst := "TestGetProductWithID"

	returnValue, err := getClient().GetProductWithID("bfbd599a-294b-4e52-989f-ec8d6b281950")

	ReturnAndError(t, tst, returnValue, err)
}

func TestUpdateProducts(t *testing.T) {
	tst := "TestUpdateProducts"

	returnValue, err := getClient().UpdateProduct("bfbd599a-294b-4e52-989f-ec8d6b281950", &UpdateProductDataBody{
		CategoryID: "27d3d24a-257b-4a88-a8eb-bb8b6121f317",
		Price:      "10.3",
		Name:       "Mert",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestCreateProducts(t *testing.T) {
	tst := "TestCreateProducts"

	returnValue, err := getClient().CreateProduct(&CreateProductDataBody{
		CategoryID: "27d3d24a-257b-4a88-a8eb-bb8b6121f317",
		Price:      "10.2",
		Name:       "Mert",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestAddProductQuestionWithID(t *testing.T) {
	tst := "TestAddProductQuestionWithID"

	returnValue, err := getClient().AddProductQuestion("bfbd599a-294b-4e52-989f-ec8d6b281950", &AddProductQuestionDataBody{
		QuestionID: "",
	})

	ReturnAndError(t, tst, returnValue, err)
}

func TestProductReorder(t *testing.T) {
	tst := "TestProductReorder"

	err := getClient().ProductReorder("bfbd599a-294b-4e52-989f-ec8d6b281950", &ProductReorderDataBody{
		SubsequentID: "",
	})

	OnlyError(t, tst, err)
}

func TestProductQuestionReorder(t *testing.T) {
	tst := "TestProductQuestionReorder"

	err := getClient().ProductQuestionReorder("bfbd599a-294b-4e52-989f-ec8d6b281950", "", &ProductQuestionReorderDataBody{
		SubsequentID: "",
	})

	OnlyError(t, tst, err)
}

func TestDeleteProduct(t *testing.T) {
	tst := "TestDeleteProduct"

	err := getClient().DeleteProduct("bfbd599a-294b-4e52-989f-ec8d6b281950")

	OnlyError(t, tst, err)
}

func TestDeleteProductQuestion(t *testing.T) {
	tst := "TestDeleteProductQuestion"

	err := getClient().DeleteProductQuestion("bfbd599a-294b-4e52-989f-ec8d6b281950", "")

	OnlyError(t, tst, err)
}
