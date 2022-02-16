package transaction

import tr "group-project1/enitities/transaction"

type Transaction interface {
	Get() ([]tr.Transactions, error)
	Insert(newTransaction tr.Transactions) (tr.Transactions, error)
	Update(trId int, newTransaction tr.Transactions) (tr.Transactions, error)
	Delete(trId int) error
}
