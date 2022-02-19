package product

import (
	"errors"
	p "group-project1/entities/product"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// ======================== Insert Product ================================
func (ur *ProductRepository) Insert(NewProduct p.Products) (p.Products, error) {
	if err := ur.db.Create(&NewProduct).Error; err != nil {
		return NewProduct, err
	}
	return NewProduct, nil
}

// ======================== Get Products ==================================
func (ur *ProductRepository) Get() ([]p.Products, error) {
	products := []p.Products{}
	ur.db.Find(&products)
	if len(products) == 0 {
		return nil, errors.New("belum ada produk yang terdaftar")
	}
	return products, nil
}

func (ur *ProductRepository) GetByID(ID uint) (p.Products, error) {
	product := p.Products{}
	if err := ur.db.Model(&product).Where("id = ?", ID).First(&product).Error; err != nil {
		return product, errors.New("belum ada produk yang terdaftar")
	}
	return product, nil
}

// ======================== Update Product ===============================
func (ur *ProductRepository) Update(UpdatedProduct p.Products) (p.Products, error) {
	res := ur.db.Model(&UpdatedProduct).Updates(UpdatedProduct)
	if res.RowsAffected == 0 {
		return UpdatedProduct, errors.New("tidak ada pemutakhiran pada data produk")
	}
	ur.db.First(&UpdatedProduct)
	return UpdatedProduct, nil
}

// ======================== Delete Product ===============================
func (ur *ProductRepository) Delete(ID uint) error {
	var product p.Products
	res := ur.db.Delete(&product, ID)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada produk yang dihapus")
	}
	return nil
}

// ============================================================================
