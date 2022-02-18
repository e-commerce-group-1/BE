package product

import (
	"group-project1/configs"
	"group-project1/entities/address"
	"group-project1/entities/order"
	"group-project1/entities/payment_method"
	"group-project1/entities/product"
	"group-project1/entities/product_category"
	"group-project1/entities/transaction"
	"group-project1/entities/user"
	ProductCatRepo "group-project1/repository/product_category"
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
	)
	db.AutoMigrate(&product.Products{})

	repo := New(db)
	productCatRepo := ProductCatRepo.New(db)
	t.Run("fail to create new product category", func(t *testing.T) {
		mockP := product.Products{
			Name:        "Sneaker 1",
			Description: "Produk 1",
			Gender:      true,
			Size:        40,
			Price:       1000000,
			Stock:       5,
		}
		_, err := repo.Insert(mockP)
		assert.NotNil(t, err)
	})

	t.Run("succeed to create new product category", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		mockP := product.Products{
			Name:              "Sneaker 1",
			Description:       "Produk 1",
			Gender:            true,
			Size:              40,
			Price:             1000000,
			Stock:             5,
			ProductCategoryID: 1,
		}
		productCatRepo.Insert(mockPC)
		res, err := repo.Insert(mockP)
		assert.Nil(t, err)
		assert.Equal(t, mockP.Name, res.Name)
		assert.Equal(t, mockP.Description, res.Description)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
	)
	db.AutoMigrate(&product.Products{})

	repo := New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("fail to get all user", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get all users", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		mockP := product.Products{
			Name:              "Sneaker 1",
			Description:       "Produk 1",
			Gender:            true,
			Size:              40,
			Price:             1000000,
			Stock:             5,
			ProductCategoryID: 1,
		}
		mockP2 := product.Products{
			Name:              "Sneaker 2",
			Description:       "Produk 2",
			Gender:            false,
			Size:              40,
			Price:             1900000,
			Stock:             2,
			ProductCategoryID: 1,
		}
		productCatRepo.Insert(mockPC)
		repo.Insert(mockP)
		repo.Insert(mockP2)

		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockP.Name, res[0].Name)
		assert.Equal(t, mockP2.Name, res[1].Name)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
	)
	db.AutoMigrate(&product.Products{})

	repo := New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("fail to update product", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		mockP := product.Products{
			Name:              "Sneaker 1",
			Description:       "Produk 1",
			Gender:            true,
			Size:              40,
			Price:             1000000,
			Stock:             5,
			ProductCategoryID: 1,
		}

		mockP2 := product.Products{
			Name:        "Sneaker 2",
			Description: "Produk 2",
		}
		productCatRepo.Insert(mockPC)
		repo.Insert(mockP)

		_, err := repo.Update(mockP2)
		assert.NotNil(t, err)
	})

	t.Run("succeed to update product", func(t *testing.T) {
		mockP2 := product.Products{
			Model:       gorm.Model{ID: 1},
			Name:        "Sneaker 2",
			Description: "Produk 2",
		}

		res, err := repo.Update(mockP2)
		assert.Nil(t, err)
		assert.Equal(t, mockP2.Name, res.Name)
		assert.Equal(t, mockP2.Description, res.Description)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
	)
	db.AutoMigrate(&product.Products{})

	repo := New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("fail to delete product", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to delete product", func(t *testing.T) {
		mockPC := product_category.ProductCategories{
			Name: "Sneakers",
		}
		mockP := product.Products{
			Name:              "Sneaker 1",
			Description:       "Produk 1",
			Gender:            true,
			Size:              40,
			Price:             1000000,
			Stock:             5,
			ProductCategoryID: 1,
		}

		productCatRepo.Insert(mockPC)
		repo.Insert(mockP)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
