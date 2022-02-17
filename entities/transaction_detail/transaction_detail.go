package transaction_detail

import (
	"gorm.io/gorm"
)

type TransactionDetails struct {
	gorm.Model
	Qty           int  `gorm:"type:int(11)"`
	Price         int  `gorm:"type:int(11)"`
	ProductID     uint `gorm:"primaryKey"`
	TransactionID uint `gorm:"primaryKey"`
}
