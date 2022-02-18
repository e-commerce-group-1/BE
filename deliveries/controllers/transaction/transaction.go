package transaction

import (
	"group-project1/deliveries/middlewares"
	"group-project1/entities/transaction"
	transactionRepo "group-project1/repository/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	Repo transactionRepo.Transaction
}

func New(transaction transactionRepo.Transaction) *TransactionController {
	return &TransactionController{
		Repo: transaction,
	}
}

func (ac *TransactionController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.Repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Menampilkan Data Transaction")
		}

		return c.JSON(http.StatusOK, TransactionResponseFormat{
			Code:    200,
			Success: true,
			Message: "Success Get Transaction",
			Data:    res,
		})
	}
}

func (ac *TransactionController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := CreateTransactionRequestFormat{}
		userId := int(middlewares.ExtractTokenUserId(c))

		if err := c.Bind(&requestFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		res, err := ac.Repo.Insert(transaction.Transactions{
			TotalPrice: requestFormat.TotalPrice,
			TotalQty:   requestFormat.TotalQty,
			UserID:     uint(userId),
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Membuat Transaction")
		}

		return c.JSON(http.StatusOK, TransactionResponseFormat{
			Code:    200,
			Success: true,
			Message: "Success Create Transaction",
			Data:    res,
		})
	}
}
