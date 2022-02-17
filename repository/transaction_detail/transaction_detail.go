package transaction_detail

import (
	td "group-project1/entities/transaction_detail"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type TransactionDetailRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionDetailRepository {
	return &TransactionDetailRepository{db: db}
}

// ======================== Insert Transaction Detail ===============================
func (ur *TransactionDetailRepository) Insert(newTransactionDetail td.TransactionDetails) (td.TransactionDetails, error) {
	if err := ur.db.Save(&newTransactionDetail).Error; err != nil {
		log.Warn("Found database error:", err)
		return newTransactionDetail, err
	}

	return newTransactionDetail, nil
}

// ======================== Get Transaction Details ==================================
func (ur *TransactionDetailRepository) Get() ([]td.TransactionDetails, error) {
	transaction_details := []td.TransactionDetails{}
	if err := ur.db.Find(&transaction_details).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return transaction_details, nil
}

// ======================== Update Transaction Detail ==============================
func (ur *TransactionDetailRepository) Update(transaction_detailId int, newTransactionDetail td.TransactionDetails) (td.TransactionDetails, error) {

	var transaction_detail td.TransactionDetails
	ur.db.First(&transaction_detail, transaction_detailId)

	if err := ur.db.Model(&transaction_detail).Updates(&newTransactionDetail).Error; err != nil {
		return transaction_detail, err
	}

	return transaction_detail, nil
}

// ======================== Delete Transaction Detail ==============================
func (ur *TransactionDetailRepository) Delete(transaction_detailId int) error {

	var transaction_detail td.TransactionDetails

	if err := ur.db.First(&transaction_detail, transaction_detailId).Error; err != nil {
		return err
	}
	ur.db.Delete(&transaction_detail, transaction_detailId)
	return nil

}

// ============================================================================
