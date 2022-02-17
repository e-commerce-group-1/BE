package address

import (
	"errors"
	a "group-project1/entities/address"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

// ======================== Insert Address ==================================
func (ur *AddressRepository) Insert(NewAddress a.Addresses) (a.Addresses, error) {
	if err := ur.db.Create(&NewAddress).Error; err != nil {
		return NewAddress, err
	}
	return NewAddress, nil
}

// ======================== Get Addresses ==================================
func (ur *AddressRepository) Get() ([]a.Addresses, error) {
	addresses := []a.Addresses{}
	ur.db.Find(&addresses)
	if len(addresses) == 0 {
		return nil, errors.New("belum ada alamat yang terdaftar")
	}
	return addresses, nil
}

// ======================== Update Address =================================
func (ur *AddressRepository) Update(UpdatedAddress a.Addresses) (a.Addresses, error) {
	res := ur.db.Model(&UpdatedAddress).Updates(UpdatedAddress)
	if res.RowsAffected == 0 {
		return UpdatedAddress, errors.New("tidak ada pemutakhiran pada data alamat")
	}
	ur.db.First(&UpdatedAddress)
	return UpdatedAddress, nil
}

// ======================== Delete Address =================================
func (ur *AddressRepository) Delete(addressId int) error {
	var address a.Addresses
	res := ur.db.Delete(&address, addressId)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada alamat yang dihapus")
	}
	return nil
}

// ============================================================================
