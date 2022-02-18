package product_category

import (
	"errors"
	pc "group-project1/entities/product_category"

	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{db: db}
}

// ======================== Insert Product Category ================================
func (ur *ProductCategoryRepository) Insert(NewProductCategory pc.ProductCategories) (pc.ProductCategories, error) {
	if err := ur.db.Create(&NewProductCategory).Error; err != nil {
		return NewProductCategory, err
	}
	return NewProductCategory, nil
}

// ======================== Get Product Categories ==================================
func (ur *ProductCategoryRepository) Get() ([]pc.ProductCategories, error) {
	product_categories := []pc.ProductCategories{}
	ur.db.Find(&product_categories)
	if len(product_categories) == 0 {
		return nil, errors.New("belum ada kategori produk yang terdaftar")
	}
	return product_categories, nil
}

// ======================== Update Product Category ===============================
func (ur *ProductCategoryRepository) Update(NewProductCategory pc.ProductCategories) (pc.ProductCategories, error) {
	res := ur.db.Model(&NewProductCategory).Update("name", NewProductCategory.Name)
	if res.RowsAffected == 0 {
		return NewProductCategory, errors.New("tidak ada pemutakhiran pada kategori produk")
	}
	ur.db.First(&NewProductCategory)
	return NewProductCategory, nil
}

// ======================== Delete Product Category ===============================
func (ur *ProductCategoryRepository) Delete(ID int) error {
	var product_category pc.ProductCategories
	res := ur.db.Delete(&product_category, ID)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada kategori produk yang dihapus")
	}
	return nil
}

// ============================================================================
