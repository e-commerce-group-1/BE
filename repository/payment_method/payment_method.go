package payment_method

import (
	pay "group-project1/entities/payment_method"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

// ======================== Insert Payment Method ================================
func (ur *PaymentMethodRepository) Insert(newPayMethod pay.PaymentMethods) (pay.PaymentMethods, error) {
	if err := ur.db.Save(&newPayMethod).Error; err != nil {
		log.Warn("Found database error:", err)
		return newPayMethod, err
	}

	return newPayMethod, nil
}

// ======================== Get Payment Methods ==================================
func (ur *PaymentMethodRepository) Get() ([]pay.PaymentMethods, error) {
	payment_methods := []pay.PaymentMethods{}
	if err := ur.db.Find(&payment_methods).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return payment_methods, nil
}

// ======================== Update Payment Method ===============================
func (ur *PaymentMethodRepository) Update(payment_methodId int, newPayMethod pay.PaymentMethods) (pay.PaymentMethods, error) {

	var payment_method pay.PaymentMethods
	ur.db.First(&payment_method, payment_methodId)

	if err := ur.db.Model(&payment_method).Updates(&newPayMethod).Error; err != nil {
		return payment_method, err
	}

	return payment_method, nil
}

// ======================== Delete Payment Method ===============================
func (ur *PaymentMethodRepository) Delete(payment_methodId int) error {

	var payment_method pay.PaymentMethods

	if err := ur.db.First(&payment_method, payment_methodId).Error; err != nil {
		return err
	}
	ur.db.Delete(&payment_method, payment_methodId)
	return nil

}

// ============================================================================
