package product

import (
	"gorm.io/gorm"

	td "group-project1/entities/transaction_detail"
)

type Products struct {
	gorm.Model
	Name                string                   `gorm:"type:varchar(100); not null"`
	Description         string                   `gorm:"type:text"`
	Gender              bool                     `gorm:"type:boolean"`
	Size                int                      `gorm:"type:int(2)"`
	Price               int                      `gorm:"type:int(8)"`
	Stock               int                      `gorm:"type:int(4)"`
	Image               string                   `gorm:"type:text"`
	Transaction_Detail  []td.Transaction_Details `gorm:"foreignKey:Product_ID"`
	Product_Category_ID uint
}
