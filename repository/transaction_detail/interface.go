package transaction_detail

import td "group-project1/enitities/transaction_detail"

type Transaction_Detail interface {
	Get() ([]td.Transaction_Detail, error)
	Insert(newTransaction_Detail td.Transaction_Detail) (td.Transaction_Detail, error)
	Update(tdId int, newTransaction_Detail td.Transaction_Detail) (td.Transaction_Detail, error)
	Delete(tdId int) error
}