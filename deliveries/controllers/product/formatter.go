package product

import "group-project1/entities/product"

// =================== Create Product ======================
type CreateProductRequestFormat struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Gender      bool   `json:"gender" form:"gender"`
	Size        uint   `json:"size" form:"size"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	Image       string `json:"image" form:"image"`
}

func (CPRF CreateProductRequestFormat) ToCreateProductEntity(ProductCategoryID uint) product.Products {
	return product.Products{
		Name:              CPRF.Name,
		Description:       CPRF.Description,
		Gender:            CPRF.Gender,
		Size:              CPRF.Size,
		Price:             CPRF.Price,
		Stock:             CPRF.Stock,
		Image:             CPRF.Image,
		ProductCategoryID: ProductCategoryID,
	}
}

type CreateProductResponseFormat struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Gender      bool   `json:"gender" form:"gender"`
	Size        uint   `json:"size" form:"size"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	Image       string `json:"image" form:"image"`
}

func ToCreateProductResponseFormat(P product.Products) CreateProductResponseFormat {
	return CreateProductResponseFormat{
		Name:        P.Name,
		Description: P.Description,
		Gender:      P.Gender,
		Size:        P.Size,
		Price:       P.Price,
		Stock:       P.Stock,
		Image:       P.Image,
	}
}

// =================== Update Product ======================
type UpdateProductRequestFormat struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Gender      bool   `json:"gender" form:"gender"`
	Size        uint   `json:"size" form:"size"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	Image       string `json:"image" form:"image"`
}

func (UPRF UpdateProductRequestFormat) ToUpdateProductEntity(ProductCategoryID uint) product.Products {
	return product.Products{
		Name:              UPRF.Name,
		Description:       UPRF.Description,
		Gender:            UPRF.Gender,
		Size:              UPRF.Size,
		Price:             UPRF.Price,
		Stock:             UPRF.Stock,
		Image:             UPRF.Image,
		ProductCategoryID: ProductCategoryID,
	}
}

type UpdateProductResponseFormat struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Gender      bool   `json:"gender"`
	Size        uint   `json:"size"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Image       string `json:"image"`
}

func ToUpdateProductResponseFormat(P product.Products) UpdateProductResponseFormat {
	return UpdateProductResponseFormat{
		Name:        P.Name,
		Description: P.Description,
		Gender:      P.Gender,
		Size:        P.Size,
		Price:       P.Price,
		Stock:       P.Stock,
		Image:       P.Image,
	}
}

// =================== Get Products =======================
type GetProductResponseFormat struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Gender      bool   `json:"gender" form:"gender"`
	Size        uint   `json:"size" form:"size"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	Image       string `json:"image" form:"image"`
}

func ToProductGetResponseFormat(Responses []product.Products) []GetProductResponseFormat {
	GetResponses := make([]GetProductResponseFormat, len(Responses))
	for i := 0; i < len(Responses); i++ {
		GetResponses[i].Name = Responses[i].Name
		GetResponses[i].Description = Responses[i].Description
		GetResponses[i].Gender = Responses[i].Gender
		GetResponses[i].Size = Responses[i].Size
		GetResponses[i].Price = Responses[i].Price
		GetResponses[i].Stock = Responses[i].Stock
		GetResponses[i].Image = Responses[i].Image
	}
	return GetResponses
}
