package transaction

import (
	"time"
)

type CreateTransactionRequestFormat struct {
	ProductID uint `json:"product_id"`
	Qty       uint `json:"qty"`
	Price     uint `json:"price"`
}

type TransactionResponseFormat struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ProductID  uint      `json:"product_id"`
	ProductQty uint      `json:"product_qty"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	Qty        uint      `json:"qty"`
	Price      uint      `json:"price"`
	Status     string    `json:"status"`
}
