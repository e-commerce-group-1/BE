package order

import (
	"group-project1/entities/address"
	o "group-project1/entities/order"
	"strings"
)

// =================== Create Order =======================
type CreateOrderRequestFormat struct {
	PaymentMethodID uint     `json:"payment_method_id " form:"payment_method_id"`
	TransactionID   []string `json:"transaction_id" form:"transaction_id"`
	Phone           string   `json:"phone" form:"phone"`
}

func (CORF CreateOrderRequestFormat) ToOrderEntity(UserID uint) o.Orders {
	return o.Orders{
		Phone:           CORF.Phone,
		PaymentMethodID: CORF.PaymentMethodID,
		UserID:          UserID,
	}
}

type OrderResponseFormat struct {
	ID              uint              `json:"id"`
	UserID          uint              `json:"user_id"`
	PaymentMethodID uint              `json:"payment_method_id"`
	TransactionID   []string          `json:"transaction_id"`
	Phone           string            `json:"phone"`
	Address         address.Addresses `json:"address"`
	Status          string            `json:"status"`
}

func ToOrderResponseFormat(order o.Orders) OrderResponseFormat {
	strArr := strings.Split(order.TransactionID, ",")
	return OrderResponseFormat{
		ID:              order.ID,
		UserID:          order.UserID,
		PaymentMethodID: order.PaymentMethodID,
		TransactionID:   strArr,
		Phone:           order.Phone,
		Address:         order.Address,
		Status:          order.Status,
	}
}

func ToOrderResponseFormatArr(orders []o.Orders) []OrderResponseFormat {
	Responses := make([]OrderResponseFormat, len(orders))
	for i := 0; i < len(orders); i++ {
		strArr := strings.Split(orders[i].TransactionID, ",")

		Responses[i].ID = orders[i].ID
		Responses[i].UserID = orders[i].UserID
		Responses[i].PaymentMethodID = orders[i].PaymentMethodID
		Responses[i].TransactionID = strArr
		Responses[i].Phone = orders[i].Phone
		Responses[i].Address = orders[i].Address
		Responses[i].Status = orders[i].Status

	}
	return Responses
}

// =================== Update Order =======================
// type UpdateOrderRequestFormat struct {
// 	Phone  string    `json:"phone" form:"phone"`
// 	Status time.Time `json:"status" form:"status"`
// }
