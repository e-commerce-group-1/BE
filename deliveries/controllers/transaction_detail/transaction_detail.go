package transaction_detail

import (
	"group-project1/entities/transaction_detail"
	trDetailRepo "group-project1/repository/transaction_detail"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionDetailController struct {
	Repo trDetailRepo.Transaction_Detail
}

func New(trDetail trDetailRepo.Transaction_Detail) *TransactionDetailController {
	return &TransactionDetailController{
		Repo: trDetail,
	}
}

func (ac *TransactionDetailController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.Repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Menampilkan Transaksi Detail")
		}

		return c.JSON(http.StatusOK, TransactionDetailResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Get Transaction Detail",
			Data: res,
		})
	}
}

func (ac *TransactionDetailController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := CreateTransactionDetailRequestFormat{}

		if err := c.Bind(&requestFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		res, err := ac.Repo.Insert(transaction_detail.TransactionDetails{
			Qty: requestFormat.Qty,
			Price: requestFormat.Price,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Membuat Transaksi Detail")
		}

		return c.JSON(http.StatusOK, TransactionDetailResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Create Transaction Detail",
			Data: res,
		})
	}
}