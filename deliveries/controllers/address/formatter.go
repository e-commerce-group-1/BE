package address

import (
	"group-project1/entities/address"
	// "gorm.io/gorm"
)

// =================== Create Address =======================
type CreateAddressRequestFormat struct {
	Street string `json:"street" form:"street"`
	City string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	ZipCode string `json:"zip_code" form:"zip_code"`
}

type CreateAddressResponseFormat struct {
	Street string `json:"street" form:"street"`
	City string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	ZipCode string `json:"zip_code" form:"zip_code"`
}

func ToCreateAddressResponseFormat(AddressResponse address.Addresses) CreateAddressResponseFormat {
	return CreateAddressResponseFormat{
		Street: AddressResponse.Street,
		City: AddressResponse.City,
		Province: AddressResponse.Province,
		ZipCode: AddressResponse.ZipCode,
	}
}

// =================== Update Address =======================
type UpdateAddressRequestFormat struct {
	Street string `json:"street" form:"street"`
	City string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	ZipCode string `json:"zip_code" form:"zip_code"`
}

func (UARF UpdateAddressRequestFormat) ToUpdateAddressRequestFormat(ID uint) address.Addresses {
	return address.Addresses{
		// Model:    gorm.Model{ID: ID},
		UserID: ID,
		Street: UARF.Street,
		City: UARF.City,
		Province: UARF.Province,
		ZipCode: UARF.ZipCode,
	}
}

type UpdateAddressResponseFormat struct {
	Street string `json:"street" form:"street"`
	City string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	ZipCode string `json:"zip_code" form:"zip_code"`
}

func ToUpdateAddressResponseFormat(AddressResponse address.Addresses) UpdateAddressResponseFormat {
	return UpdateAddressResponseFormat{
		Street: AddressResponse.Street,
		City: AddressResponse.City,
		Province: AddressResponse.Province,
		ZipCode: AddressResponse.ZipCode,
	}
}