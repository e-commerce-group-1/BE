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
	if err := ur.db.Create(&NewTransaction).Error; err != nil {
		return NewTransaction, err
	}
	return NewTransaction, nil
}

// ======================== Get Transactions ==================================
func (ur *TransactionRepository) Get() ([]t.Transactions, error) {
	transactions := []t.Transactions{}
	ur.db.Find(&transactions)
	if len(transactions) < 1 {
		return nil, errors.New("belum ada tranasaksi yang terdaftar")
	}
	return transactions, nil
}

// ======================== Update Transaction ==============================
func (ur *TransactionRepository) Update(UpdatedTransaction t.Transactions) (t.Transactions, error) {
	res := ur.db.Model(&UpdatedTransaction).Updates(UpdatedTransaction)
	if res.RowsAffected == 0 {
		return UpdatedTransaction, errors.New("tidak ada pemutakhiran pada data transaksi")
	}
	ur.db.First(&UpdatedTransaction)
	return UpdatedTransaction, nil
}

// ======================== Delete Transaction ==============================
func (ur *TransactionRepository) Delete(ID int) error {
	var transaction t.Transactions
	res := ur.db.Delete(&transaction, ID)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada user yang dihapus")
	}
	return nil

}

// ============================================================================
