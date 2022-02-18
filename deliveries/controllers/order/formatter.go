package order

import (
	"time"
)

// =================== Create Order =======================
type CreateOrderRequestFormat struct {
	Phone  string `json:"phone" form:"phone"`
	Status bool   `json:"status" form:"status"`
}

// =================== Update Order =======================
type UpdateOrderRequestFormat struct {
	Phone  string    `json:"phone" form:"phone"`
	Status time.Time `json:"status" form:"status"`
}

type OrderResponseFormat struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
