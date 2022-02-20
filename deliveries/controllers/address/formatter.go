package address

import (
	"group-project1/entities/address"
)

// =================== Create Address =======================
type CreateAddressRequestFormat struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
}

type AddressResponseFormat struct {
	ID       uint   `json:"id"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
}

func ToAddressResponseFormat(AddressResponse address.Addresses) AddressResponseFormat {
	return AddressResponseFormat{
		ID:       AddressResponse.ID,
		Street:   AddressResponse.Street,
		City:     AddressResponse.City,
		Province: AddressResponse.Province,
	}
}

// =================== Update Address =======================
type UpdateAddressRequestFormat struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
}

func (UARF UpdateAddressRequestFormat) ToUpdateAddressRequestFormat(OrderID uint) address.Addresses {
	return address.Addresses{
		OrderID:  OrderID,
		Street:   UARF.Street,
		City:     UARF.City,
		Province: UARF.Province,
	}
}

type UpdateAddressResponseFormat struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	ZipCode  string `json:"zip_code" form:"zip_code"`
}

func ToUpdateAddressResponseFormat(AddressResponse address.Addresses) UpdateAddressResponseFormat {
	return UpdateAddressResponseFormat{
		Street:   AddressResponse.Street,
		City:     AddressResponse.City,
		Province: AddressResponse.Province,
	}
}
