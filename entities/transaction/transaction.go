package transaction

import (
	td "group-project1/entities/transaction_detail"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	TotalQty          int `gorm:"type:int(11)"`
	TotalPrice        int `gorm:"type:int(11)"`
	UserID            uint
	OrderID           uint                    `gorm:"foreignKey:TransactionID"`
	TransactionDetail []td.TransactionDetails `gorm:"foreignKey:TransactionID"`
}
