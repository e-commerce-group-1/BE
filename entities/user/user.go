package user

import (
	"group-project1/entities/address"
	"group-project1/entities/transaction"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name        string                     `gorm:"type:varchar(255); not null"`
	UserName    string                     `gorm:"type:varchar(255); unique; not null"`
	Email       string                     `gorm:"type:varchar(255); unique; not null"`
	Password    string                     `gorm:"type:varchar(255); not null"`
	IsAdmin     bool                       `gorm:"type: boolean"`
	Address     address.Addresses          `gorm:"foreignKey:UserID"`
	Transaction []transaction.Transactions `gorm:"foreignKey:UserID"`
}
