package transaction

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Size      string `gorm:"type:varchar(5)"`
	Qty       uint   `gorm:"type:int(11);default:1"`
	Status    string `gorm:"type:enum('cart', 'order', 'payed', 'cancel');default:'cart'"`
	OrderID   uint
}
