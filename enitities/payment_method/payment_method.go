package payment_method

import (
	"group-project1/enitities/order"

	"gorm.io/gorm"
)

type Payment_Methods struct {
	gorm.Model
	Order []order.Orders `gorm:"foreignKey:Payment_Method_ID"`
	Name  string         `gorm:"type:varchar(255); not null"`
}
