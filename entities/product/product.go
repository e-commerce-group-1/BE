package product

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name              string `gorm:"type:varchar(100); not null"`
	Description       string `gorm:"type:text"`
	Gender            bool   `gorm:"type:boolean"`
	Size              uint
	Price             uint
	Stock             uint   `gorm:"default:0"`
	Image             string `gorm:"type:text"`
	ProductCategoryID uint
}
