package payment_method

import pay "group-project1/enitities/payment_method"

type Payment_Method interface {
	Get() ([]pay.Payment_Methods, error)
	Insert(newPayMethod pay.Payment_Methods) (pay.Payment_Methods, error)
	Update(payMethodId int, newPayMethod pay.Payment_Methods) (pay.Payment_Methods, error)
	Delete(payMethodId int) error
}
