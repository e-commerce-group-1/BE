package transaction

import (
	t "group-project1/enitities/transaction"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// ======================== Insert Transaction ===============================
func (ur *TransactionRepository) Insert(newTransaction t.Transactions) (t.Transactions, error) {
	if err := ur.db.Save(&newTransaction).Error; err != nil {
		log.Warn("Found database error:", err)
		return newTransaction, err
	}

	return newTransaction, nil
}

// ======================== Get Transactions ==================================
func (ur *TransactionRepository) Get() ([]t.Transactions, error) {
	transactions := []t.Transactions{}
	if err := ur.db.Find(&transactions).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return transactions, nil
}

// ======================== Update Transaction ==============================
func (ur *TransactionRepository) Update(transactionId int, newTransaction t.Transactions) (t.Transactions, error) {

	var transaction t.Transactions
	ur.db.First(&transaction, transactionId)

	if err := ur.db.Model(&transaction).Updates(&newTransaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// ======================== Delete Transaction ==============================
func (ur *TransactionRepository) Delete(transactionId int) error {

	var transaction t.Transactions

	if err := ur.db.First(&transaction, transactionId).Error; err != nil {
		return err
	}
	ur.db.Delete(&transaction, transactionId)
	return nil

}

// ============================================================================
