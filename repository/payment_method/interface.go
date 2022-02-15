package payment_method

import pay "group-project1/enitities/payment_method"

type Payment_Method interface {
	Get() ([]pay.Payment_Method, error)
	Insert(newPayMethod pay.Payment_Method) (pay.Payment_Method, error)
	Update(payMethodId int, newPayMethod pay.Payment_Method) (pay.Payment_Method, error)
	Delete(payMethodId int) error
}