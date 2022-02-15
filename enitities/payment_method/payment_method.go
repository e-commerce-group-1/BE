package payment_method

import (
	"gorm.io/gorm"
	"group-project1/enitities/order"

)

type Payment_Method struct {
	gorm.Model
	Order []order.Order `gorm:"foreignKey:Payment_Method_ID"`
	Name      string `gorm:"type:varchar(255); not null"`
}
