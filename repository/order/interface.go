package order

import o "group-project1/entities/order"

type Order interface {
	Insert(trxIDs []string, NewOrder o.Orders) (o.Orders, error)
	GetByUserID(UserID uint) ([]o.Orders, error)
	SetPayed(OrderID uint) (o.Orders, error)
	SetCancel(OrderID uint) (o.Orders, error)
	GetHistoryByUserID(UserID uint) ([]o.Orders, error)
}
