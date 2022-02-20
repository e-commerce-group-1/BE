package transaction

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	TrxRepo "group-project1/repository/transaction"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	Repo TrxRepo.Transaction
}

func New(repo TrxRepo.Transaction) *TransactionController {
	return &TransactionController{
		Repo: repo,
	}
}

func (tc *TransactionController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		NewTrx := CreateTransactionRequestFormat{}
		if err := c.Bind(&NewTrx); err != nil || NewTrx.Qty < 1 || NewTrx.Size == "" || NewTrx.ProductID < 1 {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := tc.Repo.Insert(NewTrx.ToTransactionEntity(uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan transaksi baru", ToTransactionResponseFormat(res)))
	}
}

func (tc *TransactionController) GetAllTrxByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)

		res, err := tc.Repo.GetAllTrxByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua transaksi berdasarkan UserID", ToTransactionResponseFormatArr(res)))
	}
}

func (tc *TransactionController) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		ProductID, _ := strconv.Atoi(c.Param("id"))
		size := Size{}
		if err := c.Bind(&size); err != nil || size.Size == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err := tc.Repo.DeleteByID(uint(ProductID), uint(UserID), size.Size)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus data transaksi", err))
	}
}

func (tc *TransactionController) FindID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		ProductID, _ := strconv.Atoi(c.Param("id"))
		size := Size{}
		if err := c.Bind(&size); err != nil || size.Size == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		res, err := tc.Repo.FindID(uint(ProductID), uint(UserID), size.Size)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan ID transaksi", res))
	}
}
