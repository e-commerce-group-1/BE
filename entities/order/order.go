package order

import (
	"group-project1/entities/address"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Phone           string            `gorm:"type:varchar(13); not null"`
	Status          string            `gorm:"type:enum('order', 'payed', 'cancel');default:'order'"`
	TransactionID   string            `gorm:"type:text"`
	Address         address.Addresses `gorm:"foreignKey:OrderID"`
	PaymentMethodID uint
	UserID          uint
}
