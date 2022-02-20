package product

import (
	"group-project1/entities/transaction"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name        string                     `gorm:"type:varchar(255);not null;unique"`
	Category    string                     `gorm:"type:enum('Sneakers','Shirts', 'Pants', 'Accessories')"`
	Description string                     `gorm:"type:text"`
	Gender      bool                       `gorm:"type:boolean"`
	Size        string                     `gorm:"type:varchar(255)"`
	Price       uint                       `gorm:"default:0"`
	Stock       uint                       `gorm:"default:0"`
	Image       string                     `gorm:"type:text"`
	Transaction []transaction.Transactions `gorm:"foreignKey:ProductID"`
}
