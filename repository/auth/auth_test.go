package auth

import (
	"group-project1/configs"
	"group-project1/entities/address"
	"group-project1/entities/order"
	"group-project1/entities/payment_method"
	"group-project1/entities/product"
	"group-project1/entities/transaction"
	"group-project1/entities/user"
	UserRepo "group-project1/repository/user"
	"group-project1/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{},
	)
	db.AutoMigrate(&address.Addresses{})

	repo := New(db)
	userRepo := UserRepo.New(db)

	t.Run("fail to login", func(t *testing.T) {
		mockUser := user.Users{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}

		_, err := repo.Login(mockUser.Email, mockUser.Password)
		assert.NotNil(t, err)
	})

	t.Run("success to login", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		userRepo.Insert(mockUser)

		res, err := repo.Login(mockUser.Email, mockUser.Password)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
		assert.Equal(t, mockUser.Email, res.Email)
		assert.Equal(t, mockUser.Password, res.Password)
	})
}
