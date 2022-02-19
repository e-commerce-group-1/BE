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
		UserID := uint(middlewares.ExtractTokenUserId(c))

		if err := c.Bind(&NewAddress); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := ac.Repo.Insert(address.Addresses{
			Street:   NewAddress.Street,
			City:     NewAddress.City,
			Province: NewAddress.Province,
			ZipCode:  NewAddress.ZipCode,
			UserID:   UserID,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan alamat baru", ToCreateAddressResponseFormat(res)))
	}
}

func (ac *AddressController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		AddID, _ := strconv.Atoi(c.Param("id"))
		var UpdateAddress = UpdateAddressRequestFormat{}

		if err := c.Bind(&UpdateAddress); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		res, err := ac.Repo.Update(UpdateAddress.ToUpdateAddressRequestFormat(uint(AddID)))

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
