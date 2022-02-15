package product_category

import pc "group-project1/enitities/product_category"

type Product_Category interface {
	Get() ([]pc.Product_Category, error)
	Insert(newProductCategory pc.Product_Category) (pc.Product_Category, error)
	Update(pcId int, newProductCategory pc.Product_Category) (pc.Product_Category, error)
	Delete(pcId int) error
}