package product

import p "group-project1/entities/product"

type Product interface {
	Get() ([]p.Products, error)
	GetByID(ID uint) (p.Products, error)
	Insert(NewProduct p.Products) (p.Products, error)
	Update(UpdatedProduct p.Products) (p.Products, error)
	Delete(ID uint) error
}
