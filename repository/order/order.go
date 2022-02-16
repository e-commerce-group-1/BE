package order

import (
	o "group-project1/enitities/order"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// ======================== Insert Order =================================
func (ur *OrderRepository) Insert(newOrder o.Order) (o.Order, error) {
	if err := ur.db.Save(&newOrder).Error; err != nil {
		log.Warn("Found database error:", err)
		return newOrder, err
	}

	return newOrder, nil
}

// ======================== Get Orders ==================================
func (ur *OrderRepository) Get() ([]o.Order, error) {
	orders := []o.Order{}
	if err := ur.db.Find(&orders).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return orders, nil
}

// ======================== Update Order ================================
func (ur *OrderRepository) Update(orderId int, newOrder o.Order) (o.Order, error) {

	var order o.Order
	ur.db.First(&order, orderId)

	if err := ur.db.Model(&order).Updates(&newOrder).Error; err != nil {
		return order, err
	}

	return order, nil
}

// ======================== Delete Order ================================
func (ur *OrderRepository) Delete(orderId int) error {

	var order o.Order

	if err := ur.db.First(&order, orderId).Error; err != nil {
		return err
	}
	ur.db.Delete(&order, orderId)
	return nil

}
// ============================================================================