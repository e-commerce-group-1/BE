package product

import (
	"group-project1/deliveries/controllers/common"
	// "group-project1/deliveries/middlewares"
	"group-project1/repository/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	repo product.Product
}

func New(repository product.Product) *ProductController {
	return &ProductController{
		repo: repository,
	}
}

func (uc *ProductController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		// isAdmin := middlewares.ExtractTokenIsAdmin(c)
		// if !isAdmin {
		// 	return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		// }
		NewProduct := CreateProductRequestFormat{}
		// ProductCategoryID, _ := strconv.Atoi(c.Param("id"))
		if err := c.Bind(&NewProduct); err != nil || NewProduct.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		res, err := uc.repo.Insert(NewProduct.ToCreateProductEntity())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan produk baru", ToCreateProductResponseFormat(res)))
	}
}

func (uc *ProductController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua produk", ToProductGetResponseFormat(res)))
	}
}

func (uc *ProductController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := uc.repo.GetByID(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan produk berdasarkan ID", ToProductGetByIDResponseFormat(res)))
	}
}

func (uc *ProductController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// isAdmin := middlewares.ExtractTokenIsAdmin(c)
		// if !isAdmin {
		// 	return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		// }
		// ProductCategoryID, _ := strconv.Atoi(c.Param("id"))
		var UpdatedProduct = UpdateProductRequestFormat{}

		if err := c.Bind(&UpdatedProduct); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(UpdatedProduct.ToUpdateProductEntity())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update produk", ToProductGetByIDResponseFormat(res)))
	}
}

func (uc *ProductController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// isAdmin := middlewares.ExtractTokenIsAdmin(c)
		// if !isAdmin {
		// 	return c.JSON(http.StatusUnauthorized, common.UnAuthorized())
		// }
		ProductCategoryID, _ := strconv.Atoi(c.Param("id"))
		err := uc.repo.Delete(uint(ProductCategoryID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus produk", err))
	}
}
