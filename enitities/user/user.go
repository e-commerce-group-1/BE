package user

import (
	"gorm.io/gorm"
	"group-project1/enitities/address"
	"group-project1/enitities/transaction"

)

type User struct {
	gorm.Model
	Address address.Address `gorm:"foreignKey:User_ID"`
	Transaction []transaction.Transaction `gorm:"foreignKey:User_ID"`
	Name      string `gorm:"type:varchar(255); not null"`
	User_name  string `gorm:"type:varchar(255); unique; not null"`
	Email     string `gorm:"type:varchar(255); unique; not null"`
	Password string `gorm:"type:varchar(255); not null"`
	Is_admin bool `gorm:"type: boolean"`
}
