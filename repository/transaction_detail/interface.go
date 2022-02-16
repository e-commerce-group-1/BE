package transaction_detail

import td "group-project1/enitities/transaction_detail"

type Transaction_Detail interface {
	Get() ([]td.Transaction_Details, error)
	Insert(newTransaction_Detail td.Transaction_Details) (td.Transaction_Details, error)
	Update(tdId int, newTransaction_Detail td.Transaction_Details) (td.Transaction_Details, error)
	Delete(tdId int) error
}
