package transaction_detail

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
	ProductRepo "group-project1/repository/product"
	ProductCatRepo "group-project1/repository/product_category"
	TrxRepo "group-project1/repository/transaction"
	UserRepo "group-project1/repository/user"
	"group-project1/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setup() *gorm.DB {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&user.Users{}, &transaction.Transactions{}, product_category.ProductCategories{}, product.Products{}, transaction_detail.TransactionDetails{})
	return db
}

func TestInsert(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
		transaction_detail.TransactionDetails{},
	)
	db.AutoMigrate(&user.Users{}, &transaction.Transactions{}, product_category.ProductCategories{}, product.Products{}, transaction_detail.TransactionDetails{})

	repo := New(db)
	userRepo := UserRepo.New(db)
	trxRepo := TrxRepo.New(db)
	productRepo := ProductRepo.New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("succeed to create transaction detail", func(t *testing.T) {
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
		productRepo.Insert(mockP)

		mockTD := transaction_detail.TransactionDetails{
			Qty:           1,
			Price:         1000000,
			ProductID:     1,
			TransactionID: 1,
		}

		res, err := repo.Insert(mockTD)
		assert.Nil(t, err)
		assert.Equal(t, mockTD.Qty, res.Qty)
		assert.Equal(t, mockTD.Price, res.Price)
		assert.Equal(t, mockTD.ProductID, res.ProductID)
		assert.Equal(t, mockTD.TransactionID, res.TransactionID)
	})

	t.Run("fail to create new transaction detail", func(t *testing.T) {
		mockTD := transaction_detail.TransactionDetails{
			Qty:           1,
			Price:         2000000,
			ProductID:     1,
			TransactionID: 2,
		}
		_, err := repo.Insert(mockTD)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	db := setup()

	repo := New(db)
	userRepo := UserRepo.New(db)
	trxRepo := TrxRepo.New(db)
	productRepo := ProductRepo.New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("fail to get all transaction detail", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get all transaction detail", func(t *testing.T) {
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
		productRepo.Insert(mockP)

		mockTD := transaction_detail.TransactionDetails{
			Qty:           1,
			Price:         1000000,
			ProductID:     1,
			TransactionID: 1,
		}

		mockTD2 := transaction_detail.TransactionDetails{
			Qty:           2,
			Price:         2000000,
			ProductID:     1,
			TransactionID: 1,
		}

		repo.Insert(mockTD)
		repo.Insert(mockTD2)

		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockTD.Qty, res[0].Qty)
		assert.Equal(t, mockTD.Price, res[0].Price)
		assert.Equal(t, mockTD.ProductID, res[0].ProductID)
		assert.Equal(t, mockTD.TransactionID, res[0].TransactionID)
		assert.Equal(t, mockTD2.Qty, res[1].Qty)
		assert.Equal(t, mockTD2.Price, res[1].Price)
		assert.Equal(t, mockTD2.ProductID, res[1].ProductID)
		assert.Equal(t, mockTD2.TransactionID, res[1].TransactionID)
	})
}

func TestUpdate(t *testing.T) {
	db := setup()

	repo := New(db)
	userRepo := UserRepo.New(db)
	trxRepo := TrxRepo.New(db)
	productRepo := ProductRepo.New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("succeed to update order", func(t *testing.T) {
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
		productRepo.Insert(mockP)

		mockTD := transaction_detail.TransactionDetails{
			Qty:           1,
			Price:         1000000,
			ProductID:     1,
			TransactionID: 1,
		}

		mockTD2 := transaction_detail.TransactionDetails{
			Model: gorm.Model{ID: 1},
			Qty:   2,
			Price: 2000000,
		}

		repo.Insert(mockTD)

		res, err := repo.Update(mockTD2)
		assert.Nil(t, err)
		assert.Equal(t, mockTD2.Qty, res.Qty)
		assert.Equal(t, mockTD2.Price, res.Price)
	})

	t.Run("fail to update order", func(t *testing.T) {
		mockTD2 := transaction_detail.TransactionDetails{
			Model: gorm.Model{ID: 3},
			Qty:   2,
			Price: 2000000,
		}

		_, err := repo.Update(mockTD2)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	db := setup()

	repo := New(db)
	userRepo := UserRepo.New(db)
	trxRepo := TrxRepo.New(db)
	productRepo := ProductRepo.New(db)
	productCatRepo := ProductCatRepo.New(db)

	t.Run("fail to delete order", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to delete order", func(t *testing.T) {
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
		productRepo.Insert(mockP)

		mockTD := transaction_detail.TransactionDetails{
			Qty:           1,
			Price:         1000000,
			ProductID:     1,
			TransactionID: 1,
		}
		repo.Insert(mockTD)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
