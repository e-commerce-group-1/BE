package order

import (
	"gorm.io/gorm"
	"group-project1/enitities/transaction"
	"time"

)

type Order struct {
	gorm.Model
	Transaction transaction.Transaction `gorm:"foreignKey:Order_ID"`
	Phone string `gorm:"type:varchar(13)"`
	Payment_Method_ID uint 
	Status    time.Time `gorm:"type:timestamp"`
}
