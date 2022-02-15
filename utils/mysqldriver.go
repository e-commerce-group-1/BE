package utils

import (
	"group-project1/configs"
	"group-project1/enitities/user"
	"group-project1/enitities/address"
	"group-project1/enitities/transaction"
	"group-project1/enitities/order"
	pay "group-project1/enitities/payment_method"
	"group-project1/enitities/product"
	p_desc "group-project1/enitities/product_description"
	p_ctg "group-project1/enitities/product_category"
	tr_detail "group-project1/enitities/transaction_detail"


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
	db.AutoMigrate(user.User{})
	db.AutoMigrate(address.Address{})
	db.AutoMigrate(transaction.Transaction{})
	db.AutoMigrate(order.Order{})
	db.AutoMigrate(pay.Payment_Method{})
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(p_desc.Product_Description{})
	db.AutoMigrate(p_ctg.Product_Category{})
	db.AutoMigrate(tr_detail.Transaction_Detail{})
}
