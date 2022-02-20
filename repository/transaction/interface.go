package transaction

import t "group-project1/entities/transaction"

type Transaction interface {
	Insert(NewTransaction t.Transactions) (t.Transactions, error)
	GetAllTrxByUserID(UserID uint) ([]t.Transactions, error)
	DeleteByID(ProductID, UserID uint) error
	FindID(ProductID, UserID uint, Size string) (uint, error)
}
