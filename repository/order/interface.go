package order

import o "group-project1/enitities/order"

type Order interface {
	Get() ([]o.Orders, error)
	Insert(newOrder o.Orders) (o.Orders, error)
	Update(orderId int, newOrder o.Orders) (o.Orders, error)
	Delete(orderId int) error
}
