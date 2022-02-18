package utils

import (
	"group-project1/configs"
	"group-project1/entities/address"
	"group-project1/entities/order"
	pay "group-project1/entities/payment_method"
	"group-project1/entities/product"
	p_ctg "group-project1/entities/product_category"
	"group-project1/entities/transaction"
	"group-project1/entities/user"

	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DropDB(config *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	Drop(db)
	return db
}

func Drop(db *gorm.DB) {
	db.Migrator().DropTable(user.Users{})
	db.Migrator().DropTable(address.Addresses{})
	db.Migrator().DropTable(transaction.Transactions{})
	db.Migrator().DropTable(order.Orders{})
	db.Migrator().DropTable(pay.PaymentMethods{})
	db.Migrator().DropTable(product.Products{})
	db.Migrator().DropTable(p_ctg.ProductCategories{})
}
