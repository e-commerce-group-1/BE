package product_category

import (
	"group-project1/entities/product"

	"gorm.io/gorm"
)

type ProductCategories struct {
	gorm.Model
	Name    string             `gorm:"type:varchar(255)"`
	Product []product.Products `gorm:"foreignKey:ProductCategoryID"`
}
