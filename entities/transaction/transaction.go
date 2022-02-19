package transaction

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	UserID    uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"primaryKey"`
	Qty       uint   `gorm:"type:int(11)"`
	Price     uint   `gorm:"type:int(11)"`
	Status    string `gorm:"type:enum('cart', 'order', 'payed', 'cancel');default:'cart'"`
	OrderID   uint   `gorm:"foreignKey:TransactionID"`
}
