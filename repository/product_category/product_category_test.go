package product_category

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
	db.AutoMigrate(&product_category.ProductCategories{})

	repo := New(db)

	t.Run("succeed to create new product category", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}

		res, err := repo.Insert(mockPC)
		assert.Nil(t, err)
		assert.Equal(t, mockPC.Name, res.Name)
	})

	t.Run("fail to create new product category", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}

		_, err := repo.Insert(mockPC)
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
	db.AutoMigrate(&product_category.ProductCategories{})

	repo := New(db)

	t.Run("fail to get all product categories", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})

	t.Run("succeed to get all product categories", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		mockPC2 := product_category.ProductCategories{
			Name: "Shirt",
		}
		repo.Insert(mockPC)
		repo.Insert(mockPC2)

		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockPC.Name, res[0].Name)
		assert.Equal(t, mockPC2.Name, res[1].Name)
	})
}
func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&product_category.ProductCategories{})

	repo := New(db)

	t.Run("succeed to update product category", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		mockPC2 := product_category.ProductCategories{
			Model: gorm.Model{ID: 1},
			Name:  "Shirt",
		}
		repo.Insert(mockPC)

		res, err := repo.Update(mockPC2)
		assert.Nil(t, err)
		assert.Equal(t, mockPC2.Name, res.Name)
	})

	t.Run("fail to update product category", func(t *testing.T) {
		mockPC2 := product_category.ProductCategories{
			Name: "Shirt",
		}

		_, err := repo.Update(mockPC2)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&product_category.ProductCategories{})

	repo := New(db)

	t.Run("fail to delete product category", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to delete product category", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		repo.Insert(mockPC)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
