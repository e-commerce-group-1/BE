package user

import (
	"group-project1/enitities/address"
	"group-project1/enitities/transaction"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name        string                     `gorm:"type:varchar(255); not null"`
	User_name   string                     `gorm:"type:varchar(255); unique; not null"`
	Email       string                     `gorm:"type:varchar(255); unique; not null"`
	Password    string                     `gorm:"type:varchar(255); not null"`
	Is_admin    bool                       `gorm:"type: boolean"`
	Address     address.Addresses          `gorm:"foreignKey:User_ID"`
	Transaction []transaction.Transactions `gorm:"foreignKey:User_ID"`
}
