package transaction

import (
	"errors"
	t "group-project1/entities/transaction"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// ======================== Insert Transaction ===============================
func (ur *TransactionRepository) Insert(NewTransaction t.Transactions) (t.Transactions, error) {
	ProductID := NewTransaction.ProductID
	UserID := NewTransaction.UserID
	Size := NewTransaction.Size
	// Bill := SetBill{}
	TrxAdded := t.Transactions{}
	// if err := ur.db.Model(&t.Transactions{}).Select("transactions.id as ID, transactions.product_id as ProductID, transactions.user_id as UserID, transactions.qty as Qty, transactions.bill as Bill, products.price as Price").Joins("inner join products on products.id = transactions.product_id").Scan(&TrxAdded); err == nil {
	// 	return TrxAdded, errors.New("produk sudah ditambahkan ke dalam keranjang")
	// }
	if err := ur.db.Where("product_id = ? AND user_id = ? AND size = ?", ProductID, UserID, Size).First(&TrxAdded).Error; err != nil {
		if NewTransaction.Qty == 0 {
			NewTransaction.Qty = 1
		}
		if err := ur.db.Create(&NewTransaction).Error; err != nil {
			return NewTransaction, err
		}
		// if err := ur.db.Table("products").Where("id = ?", ProductID).Select("price as Price").First(&Bill).Error; err != nil {
		// 	return TrxAdded, errors.New("error ketika mencari harga dari ID produk")
		// }
		// NewTransaction.Bill = bill.Price * NewTransaction.Qty

		// if err := ur.db.Table("products").Where("id = ?", ProductID).Update("stock", gorm.Expr("stock - ?", NewTransaction.Qty)).Error; err != nil {
		// 	return NewTransaction, errors.New("error ketika mengurangi stok produk")
		// }
	}
	return TrxAdded, nil
}

// func (ur *TransactionRepository) UpdateQty(ProductID uint, UserID uint, Size string, Qty uint) (t.Transactions, error) {
// 	Trx := t.Transactions{}
// 	if err := ur.db.Where("product_id = ? AND user_id = ? AND size = ?", ProductID, UserID, Size).First(&Trx).Error; err != nil {
// 		return Trx, err
// 	}

// 	if err := ur.db.Table("transactions").Where("product_id = ? AND user_id = ? AND size = ?", ProductID, UserID, Size).Update("qty", Qty).Error; err != nil {
// 		return Trx, err
// 	}

// 	// if err := ur.db.Table("products").Where("id = ?", ProductID).Update("stock", gorm.Expr("stock - ?", NewTransaction.Qty)).Error; err != nil {
// 	// 	return NewTransaction, errors.New("error ketika mengurangi stok produk")
// 	// }

// 	ur.db.Where("product_id = ? AND user_id = ? AND size = ?", ProductID, UserID, Size).First(&Trx)
// 	return Trx, nil
// }

// ======================== Get Transactions ByID ==================================
func (ur *TransactionRepository) GetAllTrxByUserID(UserID uint) ([]t.Transactions, error) {
	trx := []t.Transactions{}
	res := ur.db.Model(&t.Transactions{}).Where("user_id = ?", UserID).Find(&trx)
	if res.Error != nil {
		return trx, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return trx, nil
}

func (ur *TransactionRepository) DeleteByID(ProductID, UserID uint) error {
	var trx t.Transactions
	res := ur.db.Model(&trx).Where("product_id = ? AND user_id = ?", ProductID, UserID).Delete(&trx)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada transaksi yang dihapus")
	}
	return nil
}

func (ur *TransactionRepository) FindID(ProductID, UserID uint) (uint, error) {
	var trx t.Transactions
	if err := ur.db.Model(&trx).Where("product_id = ? AND user_id = ?", ProductID, UserID).First(&trx).Error; err != nil {
		return 0, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return trx.ID, nil
}

// ======================== Update Transaction ==============================
// func (ur *TransactionRepository) UpdateByID(ProductID uint, UserID uint, UpdatedTrx t.Transactions) (t.Transactions, error) {
// 	trxTemp := t.Transactions{}
// 	res := ur.db.Model(&trxTemp).Where("product_id = ? AND user_id = ?", ProductID, UserID).Updates(UpdatedTrx)
// 	if res.RowsAffected == 0 {
// 		return UpdatedTrx, errors.New("tidak ada pemutakhiran pada data transaksi")
// 	}
// 	ur.db.First(&UpdatedTrx)
// 	return UpdatedTrx, nil
// }

// ======================== Delete Transaction ==============================

// ======================================================= comment===========================
// import (
// 	"errors"
// 	"group-project1/entities/product"
// 	t "group-project1/entities/transaction"
// 	tr "group-project1/entities/transaction"

// 	"gorm.io/gorm"
// )

// type TransactionRepository struct {
// 	db *gorm.DB
// }

// func New(db *gorm.DB) *TransactionRepository {
// 	return &TransactionRepository{db: db}
// }

// // ======================== Insert Transaction ===============================
// func (ur *TransactionRepository) Insert(NewTransaction t.Transactions) (t.Transactions, error) {
// 	if err := ur.db.Create(&NewTransaction).Error; err != nil {
// 		return NewTransaction, err
// 	}
// 	return NewTransaction, nil
// }

// // ======================== Get Transactions ==================================
// func (ur *TransactionRepository) GetByID(ProductID uint, UserID uint) (TransactionResponseFormat, error) {
// 	trx := TransactionResponseFormat{}
// 	res := ur.db.Model(&t.Transactions{}).Where("product_id = ? AND user_id = ?", ProductID, UserID).Select("transactions.id as ID, transactions.created_at as CreatedAt, transactions.updated_at as UpdatedAt, transactions.qty as Qty, products.price as Price, products.name as Name, products.image as Image, transactions.status as Status, transactions.product_id as ProductID, products.qty as ProductQty").Joins("inner join products on products.id = transactions.product_id").Order("products.id asc").Find(&trx)
// 	if res.Error != nil {
// 		return trx, errors.New(gorm.ErrRecordNotFound.Error())
// 	}
// 	return trx, nil
// }

// // ======================== Update Transaction ==============================
// func (ur *TransactionRepository) UpdateByID(ProductID uint, UserID uint, UpdatedTrx tr.Transactions) (t.Transactions, error) {
// 	trxTemp := t.Transactions{}
// 	res := ur.db.Model(&trxTemp).Where("product_id = ? AND user_id = ?", ProductID, UserID).Find(trxTemp)
// 	if res.Error != nil {
// 		return trxTemp, nil
// 	}

// 	if _, err := ur.DeleteByID(ProductID, UserID); err != nil {
// 		return t.Transactions{}, err
// 	}

// 	ur.db.Create(&t.Transactions{UserID: UserID, ProductID: trxTemp.ProductID, Qty: trxTemp.Qty, Status: trxTemp.Status})

// 	trxTemp2 := t.Transactions{}
// 	ur.db.Model(&trxTemp2).Where("product_id = ? AND user_id = ?", trxTemp.ProductID, UserID).Find(&trxTemp2)

// 	res2 := ur.db.Model(&t.Transactions{}).Where("product_id = ? AND user_id = ?", trxTemp2.ProductID, UserID).Updates(t.Transactions{Qty: UpdatedTrx.Qty, Status: UpdatedTrx.Status}).First(&trxTemp2)
// 	if res2.Error != nil {
// 		return t.Transactions{}, res2.Error
// 	}

// 	return trxTemp2, nil
// }

// // ======================== Delete Transaction ==============================
// func (ur *TransactionRepository) DeleteByID(ProductID uint, UserID uint) (gorm.DeletedAt, error) {
// 	var transaction t.Transactions
// 	res := ur.db.Model(&transaction).Where("product_id = ? AND user_id = ?", ProductID, UserID).Delete(transaction)
// 	if res.RowsAffected == 0 {
// 		return transaction.DeletedAt, errors.New("tidak ada user yang dihapus")
// 	}
// 	return transaction.DeletedAt, nil
// }

// // ============================================================================
// func (ur *TransactionRepository) Get(UserID uint) ([]TransactionResponseFormat, error) {
// 	TrxArr := []TransactionResponseFormat{}
// 	res := ur.db.Model(&t.Transactions{}).Where("user_id = ?", UserID).Select("transactions.id as ID, transactions.created_at as CreatedAt, transactions.updated_at as UpdatedAt, transactions.qty as Qty, products.price as Price, products.name as Name, products.image as Image, transactions.status as Status, transactions.product_id as ProductID, products.qty as ProductQty").Joins("inner join products on products.id = transactions.product_id").Order("products.id asc").Find(&TrxArr)
// 	if res.Error != nil {
// 		return nil, errors.New(gorm.ErrRecordNotFound.Error())
// 	}
// 	return TrxArr, nil
// }

// func (ur *TransactionRepository) FindID(ProductID uint, UserID uint) (uint, error) {
// 	Trx := t.Transactions{}
// 	res := ur.db.Model(&Trx).Where("product_id = ? AND user_id = ?", ProductID, UserID).Find(&Trx)
// 	if res.Error != nil {
// 		return 0, errors.New(gorm.ErrRecordNotFound.Error())
// 	}
// 	return Trx.ID, nil
// }

// func (ur *TransactionRepository) InsertNew(UserID uint, NewTrx t.Transactions) (TransactionResponseFormat, error) {
// 	ProductID := NewTrx.ProductID
// 	TrxTemp := t.Transactions{}

// 	isAvailable := ur.db.Model(TrxTemp).Where("product_id = ? AND user_id = ?", ProductID, UserID).Find(&TrxTemp)
// 	if isAvailable.Error != nil {
// 		return TransactionResponseFormat{}, isAvailable.Error
// 	}

// 	if isAvailable.RowsAffected != 0 {
// 		res1, err1 := ur.GetByID(ProductID, UserID)
// 		if err1 != nil {
// 			return TransactionResponseFormat{}, err1
// 		}

// 		if _, err := ur.UpdateByID(ProductID, UserID, t.Transactions{Qty: NewTrx.Qty}); err != nil {
// 			return TransactionResponseFormat{}, err
// 		}

// 		prod := &product.Products{}
// 		res2 := ur.db.Model(&prod).Where("product_id = ?", ProductID).First(&prod)
// 		if res2.Error != nil {
// 			return TransactionResponseFormat{}, res2.Error
// 		}

// 		updateQty := ur.db.Model(&prod).Where("product_id = ?", ProductID).Update("qty", prod.Stock+(res1.Qty-NewTrx.Qty))
// 		if updateQty.Error != nil {
// 			return TransactionResponseFormat{}, updateQty.Error
// 		}

// 		res3, err3 := ur.GetByID(ProductID, UserID)
// 		if err3 != nil {
// 			return TransactionResponseFormat{}, err3
// 		}
// 		return res3, nil
// 	}

// 	NewTrx.ID = UserID
// 	if _, err := ur.Insert(NewTrx); err != nil {
// 		return TransactionResponseFormat{}, err
// 	}

// 	prod := &product.Products{}
// 	res2 := ur.db.Model(&prod).Where("product_id = ?", ProductID).First(&prod)
// 	if res2.Error != nil {
// 		return TransactionResponseFormat{}, res2.Error
// 	}

// 	if res := ur.db.Model(&product.Products{}).Where("product_id = ?", ProductID).Update("qty", prod.Stock+(res1.Qty-NewTrx.Qty)); res.Error != nil {
// 		return TransactionResponseFormat{}, res.Error
// 	}

// 	res3, err3 := ur.GetByID(ProductID, UserID)
// 	if err3 != nil {
// 		return TransactionResponseFormat{}, err3
// 	}
// 	final := TransactionResponseFormat{
// 		ID:         res3.ID,
// 		CreatedAt:  res3.CreatedAt,
// 		UpdatedAt:  res3.UpdatedAt,
// 		ProductID:  res3.ProductID,
// 		ProductQty: res3.ProductQty,
// 		Name:       res3.Name,
// 		Image:      res3.Image,
// 		Qty:        res3.Qty,
// 		Price:      res3.Qty * res3.Price,
// 		Status:     res3.Status,
// 	}
// 	return final, nil
// }

// // func (ur *TransactionRepository) DeleteNew(ProductID, UserID uint) (gorm.DeletedAt, error) {
// // 	Trx := t.Transactions{}

// // 	res1, err1 := ur.GetByID(ProductID, UserID)
// // 	if err1 != nil {
// // 		return Trx.DeletedAt, err1
// // 	}
// // 	res := ur.db.Model(&t.Transactions{}).Where("product_id = ? AND user_id = ?", ProductID, UserID).Delete(&Trx)

// // 	if res.RowsAffected == 0 {
// // 		return Trx.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
// // 	}

// // 	res2, err2 := product.New(ur.db).GetById(int(ProductID))
// // 	if err2 != nil {
// // 		return Trx.DeletedAt, err2
// // 	}

// // 	if _, err := product.New(ur.db).UpdateByIdAll(int(ProductID), templates.ProductRequest{Qty: (res2.Qty + (int(res1.Qty)))}); err != nil {
// // 		return Trx.DeletedAt, err
// // 	}

// // 	return Trx.DeletedAt, nil
// // }

// // func (ur *TransactionRepository) UpdateNew(ProductID uint, UserID uint, upCart templates.CartRequest) (templates.CartResponse, error) {
// // 	res1, err1 := ur.GetById(ProductID, UserID)
// // 	if err1 != nil {
// // 		return templates.CartResponse{}, err1
// // 	}

// // 	if _, err := ur.UpdateById(ProductID, UserID, templates.CartRequest{Qty: upCart.Qty, Status: "order"}); err != nil {
// // 		return templates.CartResponse{}, err
// // 	}

// // 	res2, err2 := product.New(ur.db).GetById(int(ProductID))

// // 	if err2 != nil {
// // 		return templates.CartResponse{}, err2
// // 	}
// // 	if _, err := product.New(ur.db).UpdateByIdAll(int(ProductID), templates.ProductRequest{Qty: (res2.Qty + (int(res1.Qty) - int(upCart.Qty)))}); err != nil {
// // 		return templates.CartResponse{}, err
// // 	}

// // 	res3, err3 := ur.GetById(ProductID, UserID)
// // 	if err3 != nil {
// // 		return templates.CartResponse{}, err3
// // 	}

// // 	res3.PriceTotal = res3.Qty * uint(res3.Price)

// // 	return res3, nil
// // }
