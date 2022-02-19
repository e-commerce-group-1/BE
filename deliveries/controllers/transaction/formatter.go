package transaction

// =================== Create Transaction =======================
type CreateTransactionRequestFormat struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Qty       uint `json:"qty" form:"total_qty"`
	Price     uint `json:"price" form:"total_price"`
}

type TransactionResponseFormat struct {
	ID         uint   `json:"id"`
	ProductID  uint   `json:"product_id"`
	ProductQty uint   `json:"qty"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Qty        uint   `json:"qty"`
	Price      uint   `json:"price"`
	Status     string `json:"status"`
}

// =================== Update Address =======================
type UpdateTransactionRequestFormat struct {
	TotalQty   int `json:"total_qty"`
	TotalPrice int `json:"total_price"`
}
