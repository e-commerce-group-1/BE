package product

import (
	p "group-project1/enitities/product"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// ======================== Insert Product ================================
func (ur *ProductRepository) Insert(newProduct p.Products) (p.Products, error) {
	if err := ur.db.Save(&newProduct).Error; err != nil {
		log.Warn("Found database error:", err)
		return newProduct, err
	}

	return newProduct, nil
}

// ======================== Get Products ==================================
func (ur *ProductRepository) Get() ([]p.Products, error) {
	products := []p.Products{}
	if err := ur.db.Find(&products).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return products, nil
}

// ======================== Update Product ===============================
func (ur *ProductRepository) Update(productId int, newProduct p.Products) (p.Products, error) {

	var product p.Products
	ur.db.First(&product, productId)

	if err := ur.db.Model(&product).Updates(&newProduct).Error; err != nil {
		return product, err
	}

	return product, nil
}

// ======================== Delete Product ===============================
func (ur *ProductRepository) Delete(productId int) error {

	var product p.Products

	if err := ur.db.First(&product, productId).Error; err != nil {
		return err
	}
	ur.db.Delete(&product, productId)
	return nil

}

// ============================================================================
