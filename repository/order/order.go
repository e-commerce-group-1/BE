package order

import (
	"errors"
	o "group-project1/entities/order"
	t "group-project1/entities/transaction"
	"strconv"
	"strings"

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
	IDArr := []int{}
	for i := 0; i < len(trxIDs); i++ {
		ID, _ := strconv.Atoi(trxIDs[i])
		IDArr = append(IDArr, ID)
	}

	if err := ur.db.Table("transactions").Where("id IN ?", IDArr).Updates(map[string]interface{}{"status": "order"}).Error; err != nil {
		return o.Orders{}, errors.New("gagal update status transaksi menjadi order")
	}

	TrxArr := []t.Transactions{}
	if err := ur.db.Table("transactions").Where("id IN ?", IDArr).Find(&TrxArr).Error; err != nil {
		return o.Orders{}, errors.New("gagal mendapatkan transaksi di dalam array ID transaksi")
	}

	for i := 0; i < len(TrxArr); i++ {
		if err := ur.db.Table("products").Where("id = ?", TrxArr[i].ProductID).Update("stock", gorm.Expr("stock - ?", TrxArr[i].Qty)).Error; err != nil {
			return o.Orders{}, errors.New("gagal update mengurangi stok produk")
		}
	}

	stringTrxID := ""
	for i := 0; i < len(trxIDs); i++ {
		stringTrxID += trxIDs[i]
		if i != len(trxIDs)-1 {
			stringTrxID += ","
		}
	}
	NewOrder.TransactionID = stringTrxID
	if err := ur.db.Create(&NewOrder).Error; err != nil {
		return NewOrder, err
	}
	return NewOrder, nil
}

// ======================== Get Orders ==================================
func (ur *OrderRepository) GetByUserID(UserID uint) ([]o.Orders, error) {
	orders := []o.Orders{}
	ur.db.Table("orders").Where("id = ?", UserID).Find(&orders)
	if len(orders) == 0 {
		return nil, errors.New("tidak terdapat order oleh UserID bersangkutan")
	}
	return orders, nil
}

// ======================== Update Order ================================
func (ur *OrderRepository) SetPayed(OrderID uint) (o.Orders, error) {
	var order o.Orders
	if err := ur.db.Model(&order).Where("order_id = ?", OrderID).Update("status", "payed").First(&order); err != nil {
		return order, errors.New("tidak terdapat order oleh UserID bersangkutan")
	}

	strArr := strings.Split(order.TransactionID, ",")
	for i := 0; i < len(strArr); i++ {
		trxID, _ := strconv.Atoi(strArr[i])
		ur.db.Table("transactions").Where("id = ?", trxID).Update("status = ?", "payed")
	}

	return order, nil
}

func (ur *OrderRepository) SetCancel(OrderID uint) (o.Orders, error) {
	var order o.Orders
	if err := ur.db.Model(&order).Where("order_id = ?", OrderID).Update("status", "cancel").First(&order); err != nil {
		return order, errors.New("tidak terdapat order oleh UserID bersangkutan")
	}

	trxIDArrStr := strings.Split(order.TransactionID, ",")
	trxIDArr := []uint{}
	for i := 0; i < len(trxIDArrStr); i++ {
		trxID, _ := strconv.Atoi(trxIDArrStr[i])
		ur.db.Table("transactions").Where("id = ?", trxID).Update("status = ?", "cancel")
		trxIDArr = append(trxIDArr, uint(trxID))
	}

	trxArr := []t.Transactions{}
	ur.db.Table("transactions").Where("id IN ?", trxIDArr).Find(&trxArr)
	for i := 0; i < len(trxArr); i++ {
		if err := ur.db.Table("products").Where("id = ?", trxArr[i].ProductID).Update("stock", gorm.Expr("stock + ?", trxArr[i].Qty)).Error; err != nil {
			return o.Orders{}, errors.New("gagal update menambahkan stok produk")
		}
	}
	return order, nil
}

func (ur *OrderRepository) GetHistoryByUserID(UserID uint) ([]o.Orders, error) {
	var orders []o.Orders
	if err := ur.db.Model(&o.Orders{}).Where("user_id = ?", UserID).Find(&orders); err != nil {
		return nil, errors.New("tidak terdapat order oleh UserID bersangkutan")
	}

	return orders, nil
}

// // ======================== Delete Order ================================
// func (ur *OrderRepository) Delete(ID int) error {
// 	var order o.Orders
// 	res := ur.db.Delete(&order, ID)
// 	if res.RowsAffected == 0 {
// 		return errors.New("tidak ada order yang dihapus")
// 	}
// 	return nil
// }

// // ============================================================================
