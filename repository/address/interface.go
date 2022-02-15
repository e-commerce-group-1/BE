package address

import a "group-project1/enitities/address"

type Address interface {
	Get() ([]a.Address, error)
	Insert(newAddress a.Address) (a.Address, error)
	Update(addressId int, newAddress a.Address) (a.Address, error)
	Delete(addressId int) error
}