package transaction

// =================== Create Transaction =======================
type CreateTransactionRequestFormat struct {
	TotalQty   int `json:"total_qty" form:"total_qty"`
	TotalPrice int `json:"total_price" form:"total_price"`
}

// =================== Update Address =======================
type UpdateTransactionRequestFormat struct {
	TotalQty   int `json:"total_qty" form:"total_qty"`
	TotalPrice int `json:"total_price" form:"total_price"`
}

type TransactionResponseFormat struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
