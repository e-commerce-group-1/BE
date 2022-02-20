package payment_method

import (
	"group-project1/configs"
	"group-project1/entities/address"
	"group-project1/entities/order"
	"group-project1/entities/payment_method"
	"group-project1/entities/product"
	"group-project1/entities/transaction"
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
		payment_method.PaymentMethods{}, product.Products{},
	)
	db.AutoMigrate(&payment_method.PaymentMethods{})

	repo := New(db)

	t.Run("succeed to create payment method", func(t *testing.T) {
		mockPM := payment_method.PaymentMethods{
			Name: "Gopay",
		}
		res, err := repo.Insert(mockPM)
		assert.Nil(t, err)
		assert.Equal(t, mockPM.Name, res.Name)
		assert.Equal(t, uint(1), res.ID)
	})

	t.Run("fail to create new user", func(t *testing.T) {
		mockPM := payment_method.PaymentMethods{
			Name: "Gopay",
		}
		_, err := repo.Insert(mockPM)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{},
	)
	db.AutoMigrate(&payment_method.PaymentMethods{})

	repo := New(db)

	t.Run("fail to get all payment methods", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})

	t.Run("succeed to get all payment methods", func(t *testing.T) {
		mockPM := payment_method.PaymentMethods{
			Name: "Gopay",
		}
		mockPM2 := payment_method.PaymentMethods{
			Name: "Dana",
		}
		repo.Insert(mockPM)
		repo.Insert(mockPM2)

		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockPM.Name, res[0].Name)
		assert.Equal(t, uint(1), res[0].ID)
		assert.Equal(t, mockPM2.Name, res[1].Name)
		assert.Equal(t, uint(2), res[1].ID)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{},
	)
	db.AutoMigrate(&payment_method.PaymentMethods{})

	repo := New(db)

	t.Run("succeed to update payment method", func(t *testing.T) {
		mockPM := payment_method.PaymentMethods{
			Name: "Gopay",
		}
		mockPM2 := payment_method.PaymentMethods{
			Model: gorm.Model{ID: 1},
			Name:  "Dana",
		}
		repo.Insert(mockPM)

		res, err := repo.Update(mockPM2)
		assert.Nil(t, err)
		assert.Equal(t, mockPM2.Name, res.Name)
	})

	t.Run("fail to update payment method", func(t *testing.T) {
		mockPM := payment_method.PaymentMethods{
			Name: "Ovo",
		}
		mockPM2 := payment_method.PaymentMethods{
			Name: "Ovo",
		}
		repo.Insert(mockPM)

		_, err := repo.Update(mockPM2)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{},
	)
	db.AutoMigrate(&payment_method.PaymentMethods{})

	repo := New(db)

	t.Run("fail to delete payment method", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to delete payment method", func(t *testing.T) {
		mockPM := payment_method.PaymentMethods{
			Name: "Gopay",
		}
		repo.Insert(mockPM)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
