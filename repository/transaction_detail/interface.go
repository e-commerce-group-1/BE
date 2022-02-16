package transaction_detail

import td "group-project1/entities/transaction_detail"

type Transaction_Detail interface {
	Get() ([]td.TransactionDetails, error)
	Insert(newTransaction_Detail td.TransactionDetails) (td.TransactionDetails, error)
	Update(tdId int, newTransaction_Detail td.TransactionDetails) (td.TransactionDetails, error)
	Delete(tdId int) error
}
