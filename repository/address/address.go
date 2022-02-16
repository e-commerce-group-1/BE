package address

import (
	a "group-project1/enitities/address"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

// ======================== Insert Address ==================================
func (ur *AddressRepository) Insert(newAddress a.Address) (a.Address, error) {
	if err := ur.db.Save(&newAddress).Error; err != nil {
		log.Warn("Found database error:", err)
		return newAddress, err
	}

	return newAddress, nil
}

// ======================== Get Addresses ==================================
func (ur *AddressRepository) Get() ([]a.Address, error) {
	addresses := []a.Address{}
	if err := ur.db.Find(&addresses).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return addresses, nil
}

// ======================== Update Address =================================
func (ur *AddressRepository) Update(addressId int, newAddress a.Address) (a.Address, error) {

	var address a.Address
	ur.db.First(&address, addressId)

	if err := ur.db.Model(&address).Updates(&newAddress).Error; err != nil {
		return address, err
	}

	return address, nil
}

// ======================== Delete Address =================================
func (ur *AddressRepository) Delete(addressId int) error {

	var address a.Address

	if err := ur.db.First(&address, addressId).Error; err != nil {
		return err
	}
	ur.db.Delete(&address, addressId)
	return nil

}
// ============================================================================