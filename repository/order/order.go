package order

import (
	"errors"
	o "group-project1/entities/order"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// ======================== Insert Order =================================
func (ur *OrderRepository) Insert(trxIDs []string, NewOrder o.Orders) (o.Orders, error) {
	IDArr = []int{}
	for i := 0; i < len(trxIDs); i++ {
		IDArr = append(IDArr, int(string(trxIDs[i])))
	}
	TrxArr := ur.db.
	if err := ur.db.Create(&NewOrder).Error; err != nil {
		return NewOrder, err
	}
	return NewOrder, nil
}

// ======================== Get Orders ==================================
func (ur *OrderRepository) Get() ([]o.Orders, error) {
	orders := []o.Orders{}
	ur.db.Find(&orders)
	if len(orders) == 0 {
		return nil, errors.New("belum ada order yang terdaftar")
	}
	return orders, nil
}

// ======================== Update Order ================================
func (ur *OrderRepository) Update(UpdatedOrder o.Orders) (o.Orders, error) {
	res := ur.db.Model(&UpdatedOrder).Updates(UpdatedOrder)
	if res.RowsAffected == 0 {
		return UpdatedOrder, errors.New("tidak ada pemutakhiran pada data order")
	}
	return UpdatedOrder, nil
}

// ======================== Delete Order ================================
func (ur *OrderRepository) Delete(ID int) error {
	var order o.Orders
	res := ur.db.Delete(&order, ID)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada order yang dihapus")
	}
	return nil
}

// ============================================================================
