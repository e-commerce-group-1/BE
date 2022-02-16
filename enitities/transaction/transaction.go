package transaction

import (
	td "group-project1/enitities/transaction_detail"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	Total_Qty          int `gorm:"type:int(11)"`
	Total_Price        int `gorm:"type:int(11)"`
	User_ID            uint
	Order_ID           uint
	Transaction_Detail []td.Transaction_Details `gorm:"foreignKey:Transaction_ID"`
}
