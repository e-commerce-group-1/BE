package product_category

import (
	"gorm.io/gorm"
	"group-project1/enitities/product"

)

type Product_Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
	Product []product.Product `gorm:"foreignKey:Product_Category_ID"`
}
