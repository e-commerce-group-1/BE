package product_description

import (
	pd "group-project1/enitities/product_description"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProductDescriptionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductDescriptionRepository {
	return &ProductDescriptionRepository{db: db}
}

// ======================== Insert Product Description ===============================
func (ur *ProductDescriptionRepository) Insert(newProductDescription pd.Product_Description) (pd.Product_Description, error) {
	if err := ur.db.Save(&newProductDescription).Error; err != nil {
		log.Warn("Found database error:", err)
		return newProductDescription, err
	}

	return newProductDescription, nil
}

// ======================== Get Product Descriptions ==================================
func (ur *ProductDescriptionRepository) Get() ([]pd.Product_Description, error) {
	product_descriptions := []pd.Product_Description{}
	if err := ur.db.Find(&product_descriptions).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return product_descriptions, nil
}

// ======================== Update Product Description ==============================
func (ur *ProductDescriptionRepository) Update(product_descriptionId int, newProductDescription pd.Product_Description) (pd.Product_Description, error) {

	var product_description pd.Product_Description
	ur.db.First(&product_description, product_descriptionId)

	if err := ur.db.Model(&product_description).Updates(&newProductDescription).Error; err != nil {
		return product_description, err
	}

	return product_description, nil
}

// ======================== Delete Product Description ==============================
func (ur *ProductDescriptionRepository) Delete(product_descriptionId int) error {

	var product_description pd.Product_Description

	if err := ur.db.First(&product_description, product_descriptionId).Error; err != nil {
		return err
	}
	ur.db.Delete(&product_description, product_descriptionId)
	return nil

}
// ============================================================================