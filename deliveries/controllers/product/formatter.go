package product

import "group-project1/entities/product"

// =================== Create Product ======================
type CreateProductRequestFormat struct {
	Name string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Gender bool `json:"gender" form:"gender"`
	Size int `json:"size" form:"size"`  
	Price int `json:"price" form:"price"`
	Stock int `json:"stock" form:"stock"`  
	Image string `json:"image" form:"image"`
}

type CreateProductResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    product.Products `json:"data"`
}

// =================== Get Products =======================
type GetProductsResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data   []product.Products `json:"data"`
}

// =================== Update Product ======================
type UpdateProductRequestFormat struct {
	Name string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Gender bool `json:"gender" form:"gender"`
	Size int `json:"size" form:"size"`  
	Price int `json:"price" form:"price"`
	Stock int `json:"stock" form:"stock"`  
	Image string `json:"image" form:"image"`
}

type UpdateProductResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    product.Products `json:"data"`
}

// =================== Delete Product ======================
type DeleteProductResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data   product.Products `json:"data"`
}