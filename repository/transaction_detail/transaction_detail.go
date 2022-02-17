package transaction_detail

import (
	"errors"
	td "group-project1/entities/transaction_detail"

	"gorm.io/gorm"
)

type TransactionDetailRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionDetailRepository {
	return &TransactionDetailRepository{db: db}
}

// ======================== Insert Transaction Detail ===============================
func (ur *TransactionDetailRepository) Insert(NewTransactionDetail td.TransactionDetails) (td.TransactionDetails, error) {
	if err := ur.db.Create(&NewTransactionDetail).Error; err != nil {
		return NewTransactionDetail, err
	}
	return NewTransactionDetail, nil
}

// ======================== Get Transaction Details ==================================
func (ur *TransactionDetailRepository) Get() ([]td.TransactionDetails, error) {
	transaction_details := []td.TransactionDetails{}
	ur.db.Find(&transaction_details)
	if len(transaction_details) < 1 {
		return nil, errors.New("belum ada detail tranasaksi yang terdaftar")
	}
	return transaction_details, nil
}

// ======================== Update Transaction Detail ==============================
func (ur *TransactionDetailRepository) Update(UpdatedTransactionDetail td.TransactionDetails) (td.TransactionDetails, error) {
	res := ur.db.Model(&UpdatedTransactionDetail).Updates(UpdatedTransactionDetail)
	if res.RowsAffected == 0 {
		return UpdatedTransactionDetail, errors.New("tidak ada pemutakhiran pada data detail transaksi")
	}
	ur.db.First(&UpdatedTransactionDetail)
	return UpdatedTransactionDetail, nil
}

// ======================== Delete Transaction Detail ==============================
func (ur *TransactionDetailRepository) Delete(ID int) error {
	var transaction_detail td.TransactionDetails
	res := ur.db.Delete(&transaction_detail, ID)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada detail transaksi yang dihapus")
	}
	return nil
}

// ============================================================================
