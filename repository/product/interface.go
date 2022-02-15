package product

import p "group-project1/enitities/product"

type Product interface {
	Get() ([]p.Product, error)
	Insert(newProduct p.Product) (p.Product, error)
	Update(productId int, newProduct p.Product) (p.Product, error)
	Delete(productId int) error
}