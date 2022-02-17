package address

import a "group-project1/entities/address"

type Address interface {
	Insert(NewAddress a.Addresses) (a.Addresses, error)
	Get() ([]a.Addresses, error)
	Update(UpdatedAddress a.Addresses) (a.Addresses, error)
	Delete(addressId int) error
}
