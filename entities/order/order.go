package order

import (
	"group-project1/entities/address"
	"group-project1/entities/transaction"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Phone           string                     `gorm:"type:varchar(13); not null"`
	Status          bool                       `gorm:"type:boolean"`
	Transaction     []transaction.Transactions `gorm:"foreignKey:OrderID"`
	Address         address.Addresses          `gorm:"foreignKey:OrderID"`
	PaymentMethodID uint
	UserID          uint
}
