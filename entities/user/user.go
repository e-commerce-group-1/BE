package user

import (
	"group-project1/entities/order"
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
	Transaction []transaction.Transactions `gorm:"foreignKey:UserID"`
	Order       []order.Orders             `gorm:"foreignKey:UserID"`
}
