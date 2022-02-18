package address

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
	UserRepo "group-project1/repository/user"
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
	db.AutoMigrate(&address.Addresses{}, &user.Users{})

	repo := New(db)
	userRepo := UserRepo.New(db)

	t.Run("succeed to create new address", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		userRepo.Insert(mockUser)

		mockAddress := address.Addresses{
			Street:   "Jl. Soedirman No.42",
			City:     "Surabaya",
			Province: "Jawa Timur",
			ZipCode:  "601111",
			UserID:   1,
		}
		res, err := repo.Insert(mockAddress)
		assert.Nil(t, err)
		assert.Equal(t, mockAddress.City, res.City)
		assert.Equal(t, mockAddress.Province, res.Province)
		assert.Equal(t, mockAddress.ZipCode, res.ZipCode)
	})

	t.Run("fail to create new address", func(t *testing.T) {
		mockAddress := address.Addresses{
			Street:   "Jl. Soedirman No.42",
			City:     "Surabaya",
			Province: "Jawa Timur",
			ZipCode:  "6011111",
		}
		_, err := repo.Insert(mockAddress)
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
	db.AutoMigrate(&address.Addresses{})

	repo := New(db)
	userRepo := UserRepo.New(db)

	t.Run("fail to get all address", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get all address", func(t *testing.T) {
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
		userRepo.Insert(mockUser)
		userRepo.Insert(mockUser2)

		mockAddress := address.Addresses{
			Street:   "Jl. Soedirman No.42",
			City:     "Surabaya",
			Province: "Jawa Timur",
			ZipCode:  "601111",
			UserID:   1,
		}
		mockAddress2 := address.Addresses{
			Street:   "Jl. Diponegoro No.42",
			City:     "Semarang",
			Province: "Jawa Tengah",
			ZipCode:  "601333",
			UserID:   2,
		}
		repo.Insert(mockAddress)
		repo.Insert(mockAddress2)

		res, err := repo.Get()

		assert.Nil(t, err)
		assert.Equal(t, mockAddress.City, res[0].City)
		assert.Equal(t, mockAddress2.City, res[1].City)
		assert.Equal(t, mockAddress.Street, res[0].Street)
		assert.Equal(t, mockAddress2.Street, res[1].Street)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&address.Addresses{})

	repo := New(db)
	userRepo := UserRepo.New(db)

	t.Run("fail update address", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		userRepo.Insert(mockUser)

		mockAddress := address.Addresses{
			Street:   "Jl. Soedirman No.42",
			City:     "Surabaya",
			Province: "Jawa Timur",
			ZipCode:  "601111",
			UserID:   1,
		}
		repo.Insert(mockAddress)

		updatedAddress := address.Addresses{
			City: "Surabaya",
		}
		_, err := repo.Update(updatedAddress)
		assert.NotNil(t, err)
	})

	t.Run("succeed to update address", func(t *testing.T) {
		mockAddress2 := address.Addresses{
			Model:  gorm.Model{ID: 1},
			Street: "Jl. Diponegoro No.99",
		}
		res, err := repo.Update(mockAddress2)

		assert.Nil(t, err)
		assert.Equal(t, "Jl. Diponegoro No.99", res.Street)
		assert.Equal(t, "Surabaya", res.City)
		assert.Equal(t, "Jawa Timur", res.Province)
		assert.Equal(t, uint(1), res.Model.ID)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&address.Addresses{})

	repo := New(db)
	userRepo := UserRepo.New(db)

	t.Run("fail to delete user", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to delete user", func(t *testing.T) {
		mockUser := user.Users{
			Name:     "Ucup",
			UserName: "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		userRepo.Insert(mockUser)

		mockAddress := address.Addresses{
			Street:   "Jl. Soedirman No.42",
			City:     "Surabaya",
			Province: "Jawa Timur",
			ZipCode:  "601111",
			UserID:   1,
		}
		repo.Insert(mockAddress)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
