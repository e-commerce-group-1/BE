package transaction_detail

import (
	"gorm.io/gorm"
)

type Transaction_Details struct {
	gorm.Model
	Qty            int  `gorm:"type:int(11)"`
	Price          int  `gorm:"type:int(11)"`
	Product_ID     uint `gorm:"primaryKey"`
	Transaction_ID uint `gorm:"primaryKey"`
}
