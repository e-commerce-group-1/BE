package payment_method

import pay "group-project1/entities/payment_method"

type Payment_Method interface {
	Get() ([]pay.PaymentMethods, error)
	Insert(newPayMethod pay.PaymentMethods) (pay.PaymentMethods, error)
	Update(payMethodId int, newPayMethod pay.PaymentMethods) (pay.PaymentMethods, error)
	Delete(payMethodId int) error
}
