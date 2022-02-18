package transaction_detail

// =================== Create Transaction Detail =======================
type CreateTransactionDetailRequestFormat struct {
	Qty   int `json:"qty" form:"qty"`
	Price int `json:"price" form:"price"`
}

// =================== Update Address =======================
type UpdateTransactionDetailRequestFormat struct {
	Qty   int `json:"qty" form:"qty"`
	Price int `json:"price" form:"price"`
}

type TransactionDetailResponseFormat struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
