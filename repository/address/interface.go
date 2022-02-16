package address

import a "group-project1/enitities/address"

type Address interface {
	Get() ([]a.Addresses, error)
	Insert(newAddress a.Addresses) (a.Addresses, error)
	Update(addressId int, newAddress a.Addresses) (a.Addresses, error)
	Delete(addressId int) error
}
