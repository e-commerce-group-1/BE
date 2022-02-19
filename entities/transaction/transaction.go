package transaction

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	UserID    uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"primaryKey"`
	Qty       int    `gorm:"type:int(11)"`
	Price     int    `gorm:"type:int(11)"`
	Status    string `gorm:"type:enum('cart', 'order', 'payed', 'cancel');default:'cart'"`
	OrderID   uint   `gorm:"foreignKey:TransactionID"`
}
