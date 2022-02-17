package payment_method

import (
	"group-project1/entities/order"

	"gorm.io/gorm"
)

type PaymentMethods struct {
	gorm.Model
	Name  string         `gorm:"type:varchar(255); not null; unique"`
	Order []order.Orders `gorm:"foreignKey:PaymentMethodID"`
}
