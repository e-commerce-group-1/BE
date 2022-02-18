package utils

import (
	"group-project1/configs"
	"group-project1/entities/address"
	"group-project1/entities/order"
	pay "group-project1/entities/payment_method"
	"group-project1/entities/product"
	p_ctg "group-project1/entities/product_category"
	"group-project1/entities/transaction"
	tr_detail "group-project1/entities/transaction_detail"
	"group-project1/entities/user"

	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.Port,
		config.Name,
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
	db.AutoMigrate(pay.PaymentMethods{})
	db.AutoMigrate(product.Products{})
	db.AutoMigrate(p_ctg.ProductCategories{})
	db.AutoMigrate(tr_detail.TransactionDetails{})
}
