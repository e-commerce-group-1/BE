package order

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Phone           string `gorm:"type:varchar(13); not null"`
	Status          bool   `gorm:"type:boolean"`
	TransactionID   uint
	PaymentMethodID uint
}
