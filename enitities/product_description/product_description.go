package product_detail

import (
	"gorm.io/gorm"
)

type Product_Description struct {
	gorm.Model
	Product_ID uint
	Name string `gorm:"type:text"`
}
