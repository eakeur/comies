package customer

import "gomies/pkg/sdk/types"

type (
	Customer struct {
		types.Entity
		Name              string
		PhoneDigest       string
		AddressCodeDigest string
	}

	Filter struct {
		Phone       string
		AddressCode string
		Name        string
	}
)

func (c Customer) Validate() error {
	return nil
}
