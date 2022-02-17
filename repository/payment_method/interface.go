package payment_method

import pay "group-project1/entities/payment_method"

type Payment_Method interface {
	Get() ([]pay.PaymentMethods, error)
	Insert(NewPayMethod pay.PaymentMethods) (pay.PaymentMethods, error)
	Update(UpdatedPaymentMethod pay.PaymentMethods) (pay.PaymentMethods, error)
	Delete(ID int) error
}
