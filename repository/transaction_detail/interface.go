package transaction_detail

import td "group-project1/entities/transaction_detail"

type Transaction_Detail interface {
	Insert(NewTransactionDetail td.TransactionDetails) (td.TransactionDetails, error)
	Get() ([]td.TransactionDetails, error)
	Update(UpdatedTransactionDetail td.TransactionDetails) (td.TransactionDetails, error)
	Delete(ID int) error
}
