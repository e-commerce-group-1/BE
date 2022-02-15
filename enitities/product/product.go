package product

import (
	"gorm.io/gorm"

	td "group-project1/enitities/transaction_detail"
	pd "group-project1/enitities/product_description"

)

type Product struct {
	gorm.Model
	Product_Category_ID uint
	Name string `gorm:"type:varchar(100); not null"`
	Gender bool `gorm:"type:boolean"`
	Size int `gorm:"type:int(2)"`
	Price int `gorm:"type:int(8)"`
	Stock int `gorm:"type:int(4)"`
	Image string `gorm:"type:text"`
	Product_Description pd.Product_Description `gorm:"foreignKey:Product_ID"`
	Transaction_Detail []td.Transaction_Detail `gorm:"foreignKey:Product_ID"`
}
