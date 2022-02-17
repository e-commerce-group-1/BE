package product_category

import (
	pc "group-project1/entities/product_category"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{db: db}
}

// ======================== Insert Product Category ================================
func (ur *ProductCategoryRepository) Insert(newProductCategory pc.ProductCategories) (pc.ProductCategories, error) {
	if err := ur.db.Save(&newProductCategory).Error; err != nil {
		log.Warn("Found database error:", err)
		return newProductCategory, err
	}

	return newProductCategory, nil
}

// ======================== Get Product Categories ==================================
func (ur *ProductCategoryRepository) Get() ([]pc.ProductCategories, error) {
	product_categories := []pc.ProductCategories{}
	if err := ur.db.Find(&product_categories).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return product_categories, nil
}

// ======================== Update Product Category ===============================
func (ur *ProductCategoryRepository) Update(product_categoryId int, newProductCategory pc.ProductCategories) (pc.ProductCategories, error) {

	var product_category pc.ProductCategories
	ur.db.First(&product_category, product_categoryId)

	if err := ur.db.Model(&product_category).Updates(&newProductCategory).Error; err != nil {
		return product_category, err
	}

	return product_category, nil
}

// ======================== Delete Product Category ===============================
func (ur *ProductCategoryRepository) Delete(product_categoryId int) error {

	var product_category pc.ProductCategories

	if err := ur.db.First(&product_category, product_categoryId).Error; err != nil {
		return err
	}
	ur.db.Delete(&product_category, product_categoryId)
	return nil

}

// ============================================================================
