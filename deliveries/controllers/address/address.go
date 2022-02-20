package address

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	"group-project1/entities/address"
	addressRepo "group-project1/repository/address"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AddressController struct {
	Repo addressRepo.Address
}

func New(address addressRepo.Address) *AddressController {
	return &AddressController{
		Repo: address,
	}
}

func (ac *AddressController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		NewAddress := CreateAddressRequestFormat{}
		OrderID := uint(middlewares.ExtractTokenUserId(c))

		if err := c.Bind(&NewAddress); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := ac.Repo.Insert(address.Addresses{
			Street:   NewAddress.Street,
			City:     NewAddress.City,
			Province: NewAddress.Province,
			OrderID:  OrderID,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan alamat baru", ToAddressResponseFormat(res)))
	}
}

func (ac *AddressController) GetByOrderID() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID, _ := strconv.Atoi(c.Param("id"))

		res, err := ac.Repo.GetByOrderID(uint(OrderID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan alamat baru", ToAddressResponseFormat(res)))
	}
}

func (ac *AddressController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID, _ := strconv.Atoi(c.Param("id"))
		var UpdateAddress = UpdateAddressRequestFormat{}

		if err := c.Bind(&UpdateAddress); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		res, err := ac.Repo.Update(UpdateAddress.ToUpdateAddressRequestFormat(uint(OrderID)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update user", ToUpdateAddressResponseFormat(res)))
	}
}

func (ac *AddressController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		AddID, _ := strconv.Atoi(c.Param("id"))

		err := ac.Repo.Delete(AddID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus alamat", err))
	}
}
