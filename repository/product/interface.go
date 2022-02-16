package product

import p "group-project1/enitities/product"

type Product interface {
	Get() ([]p.Products, error)
	Insert(newProduct p.Products) (p.Products, error)
	Update(productId int, newProduct p.Products) (p.Products, error)
	Delete(productId int) error
}
