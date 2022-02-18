package transaction

import tr "group-project1/entities/transaction"

type Transaction interface {
	Insert(NewTransaction tr.Transactions) (tr.Transactions, error)
	Get() ([]tr.Transactions, error)
	Update(UpdatedTransaction tr.Transactions) (tr.Transactions, error)
	Delete(ID int) error
}
