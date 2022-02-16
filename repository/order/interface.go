package order

import o "group-project1/enitities/order"

type Order interface {
	Get() ([]o.Order, error)
	Insert(newOrder o.Order) (o.Order, error)
	Update(orderId int, newOrder o.Order) (o.Order, error)
	Delete(orderId int) error
}