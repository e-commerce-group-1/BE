package transaction

// import (
// 	// "group-project1/deliveries/controllers/common"
// 	// "group-project1/deliveries/middlewares"
// 	"group-project1/repository/product"
// 	transactionRepo "group-project1/repository/transaction"
// 	// "net/http"

// 	// "github.com/labstack/echo/v4"
// )

// type TransactionController struct {
// 	Repo     transactionRepo.Transaction
// 	RepoProd product.Product
// }

// func New(repo transactionRepo.Transaction, repoProd product.Product) *TransactionController {
// 	return &TransactionController{
// 		Repo:     repo,
// 		RepoProd: repoProd,
// 	}
// }

// func (tc *TransactionController) Insert() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		UserID := uint(middlewares.ExtractTokenUserId(c))

// 		newTrx := CreateTransactionRequestFormat{}
// 		if err := c.Bind(&newTrx); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest())
// 		}

// 		_, err := tc.Repo.FindID(UserID, newTrx.ProductID)

// 		if err != nil {

// 		}
// 	}
// }
// func (tc *TransactionController) Get() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := tc.Repo.Get()

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Gagal Menampilkan Data Transaction")
// 		}

// 		return c.JSON(http.StatusOK, TransactionResponseFormat{
// 			Code:    200,
// 			Success: true,
// 			Message: "Success Get Transaction",
// 			Data:    res,
// 		})
// 	}
// }
