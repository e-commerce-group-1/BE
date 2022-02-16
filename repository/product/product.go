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
func (ur *ProductRepository) Insert(newProduct p.Product) (p.Product, error) {
	if err := ur.db.Save(&newProduct).Error; err != nil {
		log.Warn("Found database error:", err)
		return newProduct, err
	}

	return newProduct, nil
}

// ======================== Get Products ==================================
func (ur *ProductRepository) Get() ([]p.Product, error) {
	products := []p.Product{}
	if err := ur.db.Find(&products).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return products, nil
}

// ======================== Update Product ===============================
func (ur *ProductRepository) Update(productId int, newProduct p.Product) (p.Product, error) {

	var product p.Product
	ur.db.First(&product, productId)

	if err := ur.db.Model(&product).Updates(&newProduct).Error; err != nil {
		return product, err
	}

	return product, nil
}

// ======================== Delete Product ===============================
func (ur *ProductRepository) Delete(productId int) error {

	var product p.Product

	if err := ur.db.First(&product, productId).Error; err != nil {
		return err
	}
	ur.db.Delete(&product, productId)
	return nil

}
// ============================================================================