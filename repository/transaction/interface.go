package transaction

import tr "group-project1/enitities/transaction"

type Transaction interface {
	Get() ([]tr.Transaction, error)
	Insert(newTransaction tr.Transaction) (tr.Transaction, error)
	Update(trId int, newTransaction tr.Transaction) (tr.Transaction, error)
	Delete(trId int) error
}
