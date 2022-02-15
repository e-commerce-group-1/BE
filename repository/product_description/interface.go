package product_description

import pd "group-project1/enitities/product_description"

type Product_Description interface {
	Get() ([]pd.Product_Description, error)
	Insert(newProductDescription pd.Product_Description) (pd.Product_Description, error)
	Update(pdId int, newProductDescription pd.Product_Description) (pd.Product_Description, error)
	Delete(pdId int) error
}