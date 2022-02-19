package transaction

import (
	"errors"
	t "group-project1/entities/transaction"
	tr "group-project1/entities/transaction"

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
	if err := ur.db.Create(&NewTransaction).Error; err != nil {
		return NewTransaction, err
	}
	return NewTransaction, nil
}

// ======================== Get Transactions ==================================
func (ur *TransactionRepository) GetByID(ProductID uint, UserID uint) (TransactionResponseFormat, error) {
	trx := TransactionResponseFormat{}
	res := ur.db.Model(&t.Transactions{}).Where("product_id = ? AND user_id = ?", ProductID, UserID).Select("transactions.id as ID, transactions.created_at as CreatedAt, transactions.updated_at as UpdatedAt, transactions.qty as Qty, products.price as Price, products.name as Name, products.image as Image, transactions.status as Status, transactions.product_id as ProductID, products.qty as ProductQty").Joins("inner join products on products.id = transactions.product_id").Order("products.id asc").Find(&trx)
	if res.Error != nil {
		return trx, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return trx, nil
}

// ======================== Update Transaction ==============================
func (ur *TransactionRepository) UpdateByID(ProductID uint, UserID uint, UpdatedTrx tr.Transactions) (t.Transactions, error) {
	trxTemp := t.Transactions{}
	res := ur.db.Model(&trxTemp).Where("product_id = ? AND user_id = ?", ProductID, UserID).Find(trxTemp)
	if res.Error != nil {
		return trxTemp, nil
	}

	if _, err := ur.DeleteByID(ProductID, UserID); err != nil {
		return t.Transactions{}, err
	}

	ur.db.Create(&t.Transactions{UserID: UserID, ProductID: trxTemp.ProductID, Qty: trxTemp.Qty, Status: trxTemp.Status})

	trxTemp2 := t.Transactions{}
	ur.db.Model(&trxTemp2).Where("product_id = ? AND user_id = ?", trxTemp.ProductID, UserID).Find(&trxTemp2)

	res2 := ur.db.Model(&t.Transactions{}).Where("product_id = ? AND user_id = ?", trxTemp2.ProductID, UserID).Updates(t.Transactions{Qty: UpdatedTrx.Qty, Status: UpdatedTrx.Status}).First(&trxTemp2)
	if res2.Error != nil {
		return t.Transactions{}, res2.Error
	}

	return trxTemp2, nil
}

// ======================== Delete Transaction ==============================
func (ur *TransactionRepository) DeleteByID(ProductID uint, UserID uint) (gorm.DeletedAt, error) {
	var transaction t.Transactions
	res := ur.db.Model(&transaction).Where("product_id = ? AND user_id = ?", ProductID, UserID).Delete(transaction)
	if res.RowsAffected == 0 {
		return transaction.DeletedAt, errors.New("tidak ada user yang dihapus")
	}
	return transaction.DeletedAt, nil
}

// ============================================================================
func (ur *TransactionRepository) Get(UserID uint) ([]TransactionResponseFormat, error) {
	TrxArr := []TransactionResponseFormat{}
	res := ur.db.Model(&t.Transactions{}).Where("user_id = ?", UserID).Select("transactions.id as ID, transactions.created_at as CreatedAt, transactions.updated_at as UpdatedAt, transactions.qty as Qty, products.price as Price, products.name as Name, products.image as Image, transactions.status as Status, transactions.product_id as ProductID, products.qty as ProductQty").Joins("inner join products on products.id = transactions.product_id").Order("products.id asc").Find(&TrxArr)
	if res.Error != nil {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return TrxArr, nil
}

func (ur *TransactionRepository) FindID(ProductID uint, UserID uint) (uint, error) {
	Trx := t.Transactions{}
	res := ur.db.Model(&Trx).Where("product_id = ? AND user_id = ?", ProductID, UserID).Find(&Trx)
	if res.Error != nil {
		return 0, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return Trx.ID, nil
}
