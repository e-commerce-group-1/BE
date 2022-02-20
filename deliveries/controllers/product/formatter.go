package product

import "group-project1/entities/product"

// =================== Create Product ======================
type CreateProductRequestFormat struct {
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Gender      bool   `json:"gender" form:"gender"`
	Size        string `json:"size" form:"size"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	Image       string `json:"image" form:"image"`
}

func (CPRF CreateProductRequestFormat) ToCreateProductEntity() product.Products {
	return product.Products{
		Name:        CPRF.Name,
		Category:    CPRF.Category,
		Description: CPRF.Description,
		Gender:      CPRF.Gender,
		Size:        CPRF.Size,
		Price:       CPRF.Price,
		Stock:       CPRF.Stock,
		Image:       CPRF.Image,
	}
}

type CreateProductResponseFormat struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Gender      bool   `json:"gender"`
	Size        string `json:"size"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Image       string `json:"image"`
}

func ToCreateProductResponseFormat(P product.Products) CreateProductResponseFormat {
	return CreateProductResponseFormat{
		ID:          P.ID,
		Name:        P.Name,
		Category:    P.Category,
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
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Gender      bool   `json:"gender" form:"gender"`
	Size        string `json:"size" form:"size"`
	Price       uint   `json:"price" form:"price"`
	Stock       uint   `json:"stock" form:"stock"`
	Image       string `json:"image" form:"image"`
}

func (UPRF UpdateProductRequestFormat) ToUpdateProductEntity() product.Products {
	return product.Products{
		Name:        UPRF.Name,
		Category:    UPRF.Category,
		Description: UPRF.Description,
		Gender:      UPRF.Gender,
		Size:        UPRF.Size,
		Price:       UPRF.Price,
		Stock:       UPRF.Stock,
		Image:       UPRF.Image,
	}
}

type UpdateProductResponseFormat struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description"`
	Gender      bool   `json:"gender"`
	Size        string `json:"size"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Image       string `json:"image"`
}

func ToUpdateProductResponseFormat(P product.Products) UpdateProductResponseFormat {
	return UpdateProductResponseFormat{
		ID:          P.ID,
		Name:        P.Name,
		Category:    P.Category,
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
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gender      bool   `json:"gender"`
	Size        string `json:"size"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Image       string `json:"image"`
}

func ToProductGetResponseFormat(Responses []product.Products) []GetProductResponseFormat {
	GetResponses := make([]GetProductResponseFormat, len(Responses))
	for i := 0; i < len(Responses); i++ {
		GetResponses[i].ID = Responses[i].ID
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

func ToProductGetByIDResponseFormat(Response product.Products) GetProductResponseFormat {
	return GetProductResponseFormat{
		ID:          Response.ID,
		Name:        Response.Name,
		Description: Response.Description,
		Gender:      Response.Gender,
		Size:        Response.Size,
		Price:       Response.Price,
		Stock:       Response.Stock,
		Image:       Response.Image,
	}
}
