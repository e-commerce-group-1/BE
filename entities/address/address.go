package address

import (
	"gorm.io/gorm"
)

type Addresses struct {
	gorm.Model
	Street   string `gorm:"type:varchar(255)"`
	City     string `gorm:"type:varchar(255)"`
	Province string `gorm:"type:varchar(255)"`
	OrderID  uint
}
