package product_category

import pc "group-project1/entities/product_category"

type Product_Category interface {
	Get() ([]pc.ProductCategories, error)
	Insert(NewProductCategory pc.ProductCategories) (pc.ProductCategories, error)
	Update(NewProductCategory pc.ProductCategories) (pc.ProductCategories, error)
	Delete(ID int) error
}
