package order

import (
	// "group-project1/deliveries/controllers/auth"
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	orderRepo "group-project1/repository/order"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Repo orderRepo.Order
}

func New(order orderRepo.Order) *OrderController {
	return &OrderController{
		Repo: order,
	}
}

func (oc *OrderController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		NewOrder := CreateOrderRequestFormat{}
		if err := c.Bind(&NewOrder); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		res, err := oc.Repo.Insert(NewOrder.TransactionID, NewOrder.ToOrderEntity(uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan order baru", ToOrderResponseFormat(res)))
	}
}

func (oc *OrderController) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		res, err := oc.Repo.GetByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua order berdasarkan UserID", ToOrderResponseFormatArr(res)))
	}
}

func (oc *OrderController) SetPayed() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID, _ := strconv.Atoi(c.Param("id"))

		res, err := oc.Repo.SetPayed(uint(OrderID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status berbayar menjadi payed", ToOrderResponseFormat(res)))
	}
}

func (oc *OrderController) SetCancel() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID, _ := strconv.Atoi(c.Param("id"))

		res, err := oc.Repo.SetCancel(uint(OrderID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status berbayar menjadi payed", ToOrderResponseFormat(res)))
	}
}

func (oc *OrderController) GetHistoryByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)

		res, err := oc.Repo.GetHistoryByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan seluruh history order", ToOrderResponseFormatArr(res)))
	}
}
