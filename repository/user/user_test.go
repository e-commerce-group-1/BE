package user

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
	"group-project1/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&user.Users{})

	repo := New(db)

	t.Run("succeed to create new user", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		res, err := repo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
		assert.Equal(t, mockUser.UserName, res.UserName)
		assert.Equal(t, mockUser.Email, res.Email)
		assert.Equal(t, mockUser.Password, res.Password)
	})

	t.Run("fail to create new user", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup2",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup1232",
		}
		_, err := repo.Insert(mockUser)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&user.Users{})

	repo := New(db)

	t.Run("fail to get all user", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get all users", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		mockUser2 := user.Users{
			Name:     "Ucup2",
			UserName: "ucup2",
			Email:    "ucup2@ucup.com",
			Password: "ucup123",
		}
		repo.Insert(mockUser)
		repo.Insert(mockUser2)

		res, err := repo.Get()

		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res[0].Name)
		assert.Equal(t, mockUser2.Name, res[1].Name)
		assert.Equal(t, mockUser.UserName, res[0].UserName)
		assert.Equal(t, mockUser2.UserName, res[1].UserName)
		assert.Equal(t, mockUser.Email, res[0].Email)
		assert.Equal(t, mockUser2.Email, res[1].Email)
		assert.Equal(t, mockUser.Password, res[0].Password)
		assert.Equal(t, mockUser2.Password, res[1].Password)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&user.Users{})

	repo := New(db)

	t.Run("succeed to update users", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		mockUser2 := user.Users{
			Model:    gorm.Model{ID: 1},
			Name:     "Ucup2",
			UserName: "ucup2",
			Email:    "ucup2@ucup.com",
		}
		repo.Insert(mockUser)

		res, err := repo.Update(mockUser2)
		assert.Nil(t, err)
		assert.Equal(t, mockUser2.Name, res.Name)
		assert.Equal(t, mockUser2.UserName, res.UserName)
		assert.Equal(t, mockUser2.Email, res.Email)
		assert.Equal(t, mockUser.Password, res.Password)
	})

	t.Run("fail to update user", func(t *testing.T) {
		mockUser1 := user.Users{
			Name:     "Ucup3",
			UserName: "ucup3",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockUser2 := user.Users{
			Model:    gorm.Model{ID: 3},
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}

		repo.Insert(mockUser1)

		_, err := repo.Update(mockUser2)
		assert.NotNil(t, err)
	})
}
