package utils

import (
	"group-project1/configs"
	"group-project1/enitities/address"
	"group-project1/enitities/order"
	pay "group-project1/enitities/payment_method"
	"group-project1/enitities/product"
	p_ctg "group-project1/enitities/product_category"
	"group-project1/enitities/transaction"
	tr_detail "group-project1/enitities/transaction_detail"
	"group-project1/enitities/user"

	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(user.Users{})
	db.AutoMigrate(address.Addresses{})
	db.AutoMigrate(transaction.Transactions{})
	db.AutoMigrate(order.Orders{})
	db.AutoMigrate(pay.Payment_Methods{})
	db.AutoMigrate(product.Products{})
	db.AutoMigrate(p_ctg.Product_Categories{})
	db.AutoMigrate(tr_detail.Transaction_Details{})
}
