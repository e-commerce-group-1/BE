package payment_method

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	PMEntity "group-project1/entities/payment_method"
	PM "group-project1/repository/payment_method"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PMController struct {
	repo PM.Payment_Method
}

func New(repository PM.Payment_Method) *PMController {
	return &PMController{
		repo: repository,
	}
}

func (uc *PMController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}

		NewPM := CreatePMRequestFormat{}

		if err := c.Bind(&NewPM); err != nil || NewPM.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		newPM := PMEntity.PaymentMethods{
			Name: NewPM.Name,
		}
		res, err := uc.repo.Insert(newPM)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan kategori payment method baru", ToCreatePMResponseFormat(res)))
	}
}

func (uc *PMController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}
		PMID, _ := strconv.Atoi((c.Param("id")))
		var UpdatedPM = UpdatePMRequestFormat{}

		if err := c.Bind(&UpdatedPM); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(UpdatedPM.ToUpdatePMRequestFormat(uint(PMID)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update payment method baru", ToUpdatePMResponseFormat(res)))
	}
}

func (uc *PMController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}

		res, err := uc.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua payment method", ToPMGetResponseFormat(res)))
	}
}

func (uc *PMController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}
		PMID, _ := strconv.Atoi((c.Param("id")))
		err := uc.repo.Delete(PMID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus payment method", err))
	}
}
