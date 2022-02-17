package order

import (
	// "group-project1/deliveries/controllers/auth"
	"group-project1/entities/order"
	orderRepo "group-project1/repository/order"
	"net/http"

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

func (ac *OrderController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.Repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Menampilkan Data Orders")
		}

		return c.JSON(http.StatusOK, OrderResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Get Orders",
			Data: res,
		})
	}
}

func (ac *OrderController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := CreateOrderRequestFormat{}

		if err := c.Bind(&requestFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		res, err := ac.Repo.Insert(order.Orders{
			Phone: requestFormat.Phone,
			Status: requestFormat.Status,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Memasukkan Data Alamat")
		}

		return c.JSON(http.StatusOK, OrderResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Create Orders",
			Data: res,
		})
	}
}