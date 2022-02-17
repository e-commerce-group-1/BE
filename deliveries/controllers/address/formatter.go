package address

import "group-project1/entities/address"

// =================== Create Address =======================
type CreateAddressRequestFormat struct {
	Street string `json:"street" form:"street"`
	City string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	Zipcode string `json:"zipcode" form:"zipcode"`
}

type CreateAddressResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    address.Addresses `json:"data"`
}

// =================== Get Addresses =======================
type GetAddressesResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data   []address.Addresses `json:"data"`
}

// =================== Update Address =======================
type UpdateAddressRequestFormat struct {
	Street string `json:"street" form:"street"`
	City string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	Zipcode string `json:"zipcode" form:"zipcode"`
}

type UpdateAddressResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data    address.Addresses `json:"data"`
}

// =================== Delete Address =======================
type DeleteAddressResponseFormat struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string        `json:"message"`
	Data   address.Addresses `json:"data"`
}