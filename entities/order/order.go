package order

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Phone           string    `gorm:"type:varchar(13)"`
	Status          time.Time `gorm:"type:timestamp"`
	TransactionID   uint
	PaymentMethodID uint
}
