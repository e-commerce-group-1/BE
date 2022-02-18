package transaction

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	TotalQty   int    `gorm:"type:int(11)"`
	TotalPrice int    `gorm:"type:int(11)"`
	Status     string `gorm:"type:enum('cart', 'order', 'payed', 'cancel');default:'cart'"`
	OrderID    uint   `gorm:"foreignKey:TransactionID"`
	UserID     uint
}
