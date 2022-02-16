package order

import (
	"group-project1/enitities/transaction"
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Transaction       transaction.Transactions `gorm:"foreignKey:Order_ID"`
	Phone             string                   `gorm:"type:varchar(13)"`
	Status            time.Time                `gorm:"type:timestamp"`
	Payment_Method_ID uint
}
