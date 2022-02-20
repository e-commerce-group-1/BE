package transaction

import t "group-project1/entities/transaction"

type Transaction interface {
	Insert(NewTransaction t.Transactions) (t.Transactions, error)
	GetAllTrxByUserID(UserID uint) ([]t.Transactions, error)
	DeleteByID(ProductID uint, UserID uint) error
	FindID(UserID uint, ProductID uint) (uint, error)
}
