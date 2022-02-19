package transaction

import tr "group-project1/entities/transaction"

type Transaction interface {
	Insert(NewTrx tr.Transactions) (tr.Transactions, error)
	GetByID(ProductID uint, UserID uint) (tr.Transactions, error)
	UpdateByID(ProductID uint, UserID uint, UpdatedTrx tr.Transactions) (tr.Transactions, error)
	DeleteByID(ProductID uint, UserID uint) error
	Get(UserID uint) ([]tr.Transactions, error)
	FindID(UserID uint, ProductID uint) (uint, error)
}
