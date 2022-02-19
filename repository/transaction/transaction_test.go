// package transaction

// import (
// 	"group-project1/configs"
// 	"group-project1/entities/address"
// 	"group-project1/entities/order"
// 	"group-project1/entities/payment_method"
// 	"group-project1/entities/product"
// 	"group-project1/entities/product_category"
// 	"group-project1/entities/transaction"
// 	"group-project1/entities/user"
// 	UserRepo "group-project1/repository/user"
// 	"group-project1/utils"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )

// func TestInsert(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
// 		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
// 	)
// 	db.AutoMigrate(&user.Users{}, &transaction.Transactions{})

// 	repo := New(db)
// 	userRepo := UserRepo.New(db)

// 	t.Run("succeed to create new transaction", func(t *testing.T) {
// 		mockUser := user.Users{
// 			Name:     "Ucup",
// 			UserName: "ucup",
// 			Email:    "ucup@ucup.com",
// 			Password: "ucup123",
// 		}
// 		userRepo.Insert(mockUser)

// 		mockT := transaction.Transactions{
// 			TotalQty:   3,
// 			TotalPrice: 4500000,
// 			UserID:     1,
// 		}

// 		res, err := repo.Insert(mockT)
// 		assert.Nil(t, err)
// 		assert.Equal(t, mockT.TotalQty, res.TotalQty)
// 		assert.Equal(t, mockT.TotalPrice, res.TotalPrice)
// 	})

// 	t.Run("fail to create new transaction", func(t *testing.T) {
// 		mockT := transaction.Transactions{
// 			Model:      gorm.Model{ID: 1},
// 			TotalQty:   3,
// 			TotalPrice: 4500000,
// 			UserID:     1,
// 		}
// 		_, err := repo.Insert(mockT)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestGet(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
// 		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
// 	)
// 	db.AutoMigrate(&user.Users{}, &transaction.Transactions{})

// 	repo := New(db)
// 	userRepo := UserRepo.New(db)

// 	t.Run("fail to get all transactions", func(t *testing.T) {
// 		res, err := repo.Get()
// 		assert.Nil(t, res)
// 		assert.NotNil(t, err)
// 	})

// 	t.Run("succeed to get all transactions", func(t *testing.T) {
// 		mockUser := user.Users{
// 			Name:     "Ucup",
// 			UserName: "ucup",
// 			Email:    "ucup@ucup.com",
// 			Password: "ucup123",
// 		}
// 		userRepo.Insert(mockUser)

// 		mockT := transaction.Transactions{
// 			TotalQty:   3,
// 			TotalPrice: 4500000,
// 			UserID:     1,
// 		}
// 		mockT2 := transaction.Transactions{
// 			TotalQty:   4,
// 			TotalPrice: 4500000,
// 			UserID:     1,
// 		}

// 		repo.Insert(mockT)
// 		repo.Insert(mockT2)
// 		res, err := repo.Get()

// 		assert.Nil(t, err)
// 		assert.Equal(t, mockT.TotalPrice, res[0].TotalPrice)
// 		assert.Equal(t, mockT.TotalQty, res[0].TotalQty)
// 		assert.Equal(t, mockT2.TotalPrice, res[1].TotalPrice)
// 		assert.Equal(t, mockT2.TotalQty, res[1].TotalQty)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
// 		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
// 	)
// 	db.AutoMigrate(&user.Users{}, &transaction.Transactions{})

// 	repo := New(db)
// 	userRepo := UserRepo.New(db)

// 	t.Run("succeed to update transaction", func(t *testing.T) {
// 		mockUser := user.Users{
// 			Name:     "Ucup",
// 			UserName: "ucup",
// 			Email:    "ucup@ucup.com",
// 			Password: "ucup123",
// 		}
// 		userRepo.Insert(mockUser)

// 		mockT := transaction.Transactions{
// 			TotalQty:   3,
// 			TotalPrice: 4500000,
// 			UserID:     1,
// 		}
// 		mockT2 := transaction.Transactions{
// 			Model:      gorm.Model{ID: 1},
// 			TotalQty:   1,
// 			TotalPrice: 1000000,
// 			UserID:     1,
// 		}

// 		repo.Insert(mockT)

// 		res, err := repo.Update(mockT2)
// 		assert.Nil(t, err)
// 		assert.Equal(t, mockT2.TotalPrice, res.TotalPrice)
// 	})

// 	t.Run("fail to update transaction", func(t *testing.T) {
// 		mockT2 := transaction.Transactions{
// 			TotalQty:   1,
// 			TotalPrice: 1000000,
// 			UserID:     1,
// 		}

// 		_, err := repo.Update(mockT2)

// 		assert.NotNil(t, err)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	db.Migrator().DropTable(user.Users{}, address.Addresses{}, transaction.Transactions{}, order.Orders{},
// 		payment_method.PaymentMethods{}, product.Products{}, product_category.ProductCategories{},
// 	)
// 	db.AutoMigrate(&user.Users{}, &transaction.Transactions{})

// 	repo := New(db)
// 	userRepo := UserRepo.New(db)

// 	t.Run("fail to delete user", func(t *testing.T) {
// 		err := repo.Delete(1)
// 		assert.NotNil(t, err)
// 	})

// 	t.Run("succeed to delete user", func(t *testing.T) {
// 		mockUser := user.Users{
// 			Name:     "Ucup",
// 			UserName: "ucup",
// 			Email:    "ucup@ucup.com",
// 			Password: "ucup123",
// 		}
// 		userRepo.Insert(mockUser)

// 		mockT := transaction.Transactions{
// 			TotalQty:   3,
// 			TotalPrice: 4500000,
// 			UserID:     1,
// 		}
// 		repo.Insert(mockT)

// 		err := repo.Delete(1)
// 		assert.Nil(t, err)
// 	})
// }
