package order

import (
	"group-project1/configs"
	"group-project1/entities/address"
	"group-project1/entities/order"
	"group-project1/entities/payment_method"
	"group-project1/entities/product"
	"group-project1/entities/product_category"
	"group-project1/entities/transaction"
	"group-project1/entities/transaction_detail"
	"group-project1/entities/user"
	PM "group-project1/repository/payment_method"
	TrxRepo "group-project1/repository/transaction"
	UserRepo "group-project1/repository/user"
	"group-project1/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&order.Orders{}, &user.Users{}, &transaction.Transactions{})

	repo := New(db)
	userRepo := UserRepo.New(db)
	trxRepo := TrxRepo.New(db)
	pmRepo := PM.New(db)

	t.Run("succeed to create new order", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		userRepo.Insert(mockUser)

		mockT := transaction.Transactions{
			TotalQty:   3,
			TotalPrice: 4500000,
			UserID:     1,
		}
		trxRepo.Insert(mockT)

		mockPM := payment_method.PaymentMethods{
			Name: "Gopay",
		}
		pmRepo.Insert(mockPM)

		mockO := order.Orders{
			Phone:           "0812345678910",
			TransactionID:   1,
			PaymentMethodID: 1,
		}

		res, err := repo.Insert(mockO)
		assert.Nil(t, err)
		assert.Equal(t, mockO.Phone, res.Phone)
	})

	t.Run("fail to create new order", func(t *testing.T) {
		mockO := order.Orders{}
		_, err := repo.Insert(mockO)
		assert.NotNil(t, err)
	})
}
