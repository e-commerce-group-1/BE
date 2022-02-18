package product_category

import (
	"group-project1/deliveries/controllers/common"
	"group-project1/deliveries/middlewares"
	"group-project1/entities/product_category"
	PC "group-project1/repository/product_category"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProdCatController struct {
	repo PC.ProductCategoryRepository
}

func New(repository PC.ProductCategoryRepository) *ProdCatController {
	return &ProdCatController{
		repo: repository,
	}
}

func (uc *ProdCatController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}

		NewPC := CreateProdCatRequestFormat{}

		if err := c.Bind(&NewPC); err != nil || NewPC.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		newPC := product_category.ProductCategories{
			Name: NewPC.Name,
		}
		res, err := uc.repo.Insert(newPC)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan kategori produk baru", ToCreateUserResponseFormat(res)))
	}
}

func (uc *ProdCatController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}
		ProdCatID, _ := strconv.Atoi((c.Param("id")))
		var UpdatedPC = UpdateProdCatRequestFormat{}

		if err := c.Bind(&UpdatedPC); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(UpdatedPC.ToUpdateProdCatRequestFormat(uint(ProdCatID)))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update kategori produk", ToUpdateProdCatResponseFormat(res)))
	}
}

func (uc *ProdCatController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}

		res, err := uc.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua kategori produk", ToProdCatGetResponseFormat(res)))
	}
}

func (uc *ProdCatController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		}
		ProdCatID, _ := strconv.Atoi((c.Param("id")))
		err := uc.repo.Delete(ProdCatID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus kategori produk", err))
	}
}
