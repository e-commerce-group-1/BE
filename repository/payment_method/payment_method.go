package payment_method

import (
	"errors"
	pay "group-project1/entities/payment_method"

	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

// ======================== Insert Payment Method ================================
func (ur *PaymentMethodRepository) Insert(NewPayMethod pay.PaymentMethods) (pay.PaymentMethods, error) {
	if err := ur.db.Create(&NewPayMethod).Error; err != nil {
		return NewPayMethod, err
	}
	return NewPayMethod, nil
}

// ======================== Get Payment Methods ==================================
func (ur *PaymentMethodRepository) Get() ([]pay.PaymentMethods, error) {
	payment_methods := []pay.PaymentMethods{}
	ur.db.Find(&payment_methods)
	if len(payment_methods) == 0 {
		return nil, errors.New("belum ada payment method yang terdaftar")
	}
	return payment_methods, nil
}

// ======================== Update Payment Method ===============================
func (ur *PaymentMethodRepository) Update(UpdatedPaymentMethod pay.PaymentMethods) (pay.PaymentMethods, error) {
	res := ur.db.Model(&UpdatedPaymentMethod).Update("name", UpdatedPaymentMethod.Name)
	if res.RowsAffected == 0 {
		return UpdatedPaymentMethod, errors.New("tidak ada pemutakhiran pada payment method")
	}
	ur.db.First(&UpdatedPaymentMethod)
	return UpdatedPaymentMethod, nil
}

// ======================== Delete Payment Method ===============================
func (ur *PaymentMethodRepository) Delete(ID int) error {
	var payment_method pay.PaymentMethods
	res := ur.db.Delete(&payment_method, ID)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada payment method yang dihapus")
	}
	return nil
}

// ============================================================================
