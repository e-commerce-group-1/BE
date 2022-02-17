package order

import (
	"group-project1/entities/transaction"
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Transaction     transaction.Transactions `gorm:"foreignKey:OrderID"`
	Phone           string                   `gorm:"type:varchar(13)"`
	Status          time.Time                `gorm:"type:timestamp"`
	PaymentMethodID uint
}
