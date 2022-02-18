package product_category

import (
	"group-project1/entities/product_category"

	"gorm.io/gorm"
)

// =================== Create Product Category =======================
type CreateProdCatRequestFormat struct {
	Name string `json:"name" form:"name"`
}

type CreateProdCatResponseFormat struct {
	Name string `json:"name"`
}

func ToCreateUserResponseFormat(ProdCatResponse product_category.ProductCategories) CreateProdCatResponseFormat {
	return CreateProdCatResponseFormat{
		Name: ProdCatResponse.Name,
	}
}

// =================== Update ProdCat =======================
type UpdateProdCatRequestFormat struct {
	Name string `json:"name" form:"name"`
}

func (UURF UpdateProdCatRequestFormat) ToUpdateProdCatRequestFormat(ID uint) product_category.ProductCategories {
	return product_category.ProductCategories{
		Model: gorm.Model{ID: ID},
		Name:  UURF.Name,
	}
}

type UpdateProdCatResponseFormat struct {
	Name string `json:"name"`
}

func ToUpdateProdCatResponseFormat(ProdCatResponse product_category.ProductCategories) UpdateProdCatResponseFormat {
	return UpdateProdCatResponseFormat{
		Name: ProdCatResponse.Name,
	}
}

type ProdCatGetResponseFormat struct {
	Name string `json:"name"`
}

func ToProdCatGetResponseFormat(ProdCatResponses []product_category.ProductCategories) []ProdCatGetResponseFormat {
	ProdCatGetResponses := make([]ProdCatGetResponseFormat, len(ProdCatResponses))
	for i := 0; i < len(ProdCatResponses); i++ {
		ProdCatGetResponses[i].Name = ProdCatResponses[i].Name
	}
	return ProdCatGetResponses
}
