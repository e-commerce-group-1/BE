package address

import (
	"group-project1/deliveries/controllers/auth"
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

func (ac *AddressController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.Repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Menampilkan Data Alamat")
		}

		return c.JSON(http.StatusOK, GetAddressesResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Get Addresses",
			Data: res,
		})
	}
}

func (ac *AddressController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := CreateAddressRequestFormat{}
		userId := int(auth.ExtractTokenUserId(c))

		if err := c.Bind(&requestFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		res, err := ac.Repo.Insert(address.Addresses{
			Street: requestFormat.Street,
			City: requestFormat.City,
			Province: requestFormat.Province,
			Zipcode: requestFormat.Zipcode,
			User_ID: uint(userId),
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Gagal Memasukkan Data Alamat")
		}

		return c.JSON(http.StatusOK, CreateAddressResponseFormat{
			Code: 200,
			Success: true,
			Message: "Success Create Address",
			Data: res,
		})
	}
}

func (ac *AddressController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newAddress = UpdateAddressRequestFormat{}
		addressId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newAddress); err != nil {
			return c.JSON(http.StatusBadRequest, "Ada yang salah dengan input")
		}

		res, err := ac.Repo.Update(addressId, address.Task{Nama: newTask.Nama, Priority: newTask.Priority})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Task", res))
	}
}

func (ac *AddressController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, _ := strconv.Atoi(c.Param("id"))

		err := ac.Repo.Delete(taskId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Task", nil))
	}
}