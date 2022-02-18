package order

import o "group-project1/entities/order"

type Order interface {
	Insert(NewOrder o.Orders) (o.Orders, error)
	Get() ([]o.Orders, error)
	Update(UpdatedOrder o.Orders) (o.Orders, error)
	Delete(ID int) error
}
