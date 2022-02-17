package product_category

import pc "group-project1/entities/product_category"

type Product_Category interface {
	Get() ([]pc.ProductCategories, error)
	Insert(newProductCategory pc.ProductCategories) (pc.ProductCategories, error)
	Update(pcId int, newProductCategory pc.ProductCategories) (pc.ProductCategories, error)
	Delete(pcId int) error
}
