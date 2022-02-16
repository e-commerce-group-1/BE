package product_category

import pc "group-project1/enitities/product_category"

type Product_Category interface {
	Get() ([]pc.Product_Categories, error)
	Insert(newProductCategory pc.Product_Categories) (pc.Product_Categories, error)
	Update(pcId int, newProductCategory pc.Product_Categories) (pc.Product_Categories, error)
	Delete(pcId int) error
}
