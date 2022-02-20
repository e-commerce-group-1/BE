package transaction

import t "group-project1/entities/transaction"

// =================== Create Transaction =======================
type CreateTransactionRequestFormat struct {
	Qty       uint   `json:"qty" form:"total_qty"`
	Size      string `json:"size" form:"size"`
	ProductID uint   `json:"product_id" form:"product_id"`
	UserID    uint
}

func (CTRF CreateTransactionRequestFormat) ToTransactionEntity(UserID uint) t.Transactions {
	return t.Transactions{
		ProductID: CTRF.ProductID,
		UserID:    UserID,
		Size:      CTRF.Size,
		Qty:       CTRF.Qty,
	}
}

// func ToCreateResponse(NewTrx t.Transactions)

type TransactionResponseFormat struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	Size      string `json:"size"`
	Qty       uint   `json:"qty"`
	Status    string `json:"status"`
}

func ToTransactionResponseFormat(trx t.Transactions) TransactionResponseFormat {
	return TransactionResponseFormat{
		ID:        trx.ID,
		ProductID: trx.ProductID,
		UserID:    trx.UserID,
		Size:      trx.Size,
		Qty:       trx.Qty,
		Status:    trx.Status,
	}
}

func ToTransactionResponseFormatArr(trx []t.Transactions) []TransactionResponseFormat {
	ResArr := make([]TransactionResponseFormat, len(trx))
	for i := 0; i < len(trx); i++ {
		ResArr[i].ID = trx[i].ID
		ResArr[i].ProductID = trx[i].ProductID
		ResArr[i].UserID = trx[i].UserID
		ResArr[i].Size = trx[i].Size
		ResArr[i].Qty = trx[i].Qty
		ResArr[i].Status = trx[i].Status
	}
	return ResArr
}

// =================== Update Address =======================
type UpdateTransactionRequestFormat struct {
	TotalQty   int `json:"total_qty"`
	TotalPrice int `json:"total_price"`
}
