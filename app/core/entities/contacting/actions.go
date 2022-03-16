package contacting

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {

	// ListAddresses retrieves all the addresses of a specific target id
	ListAddresses(context.Context, id.External) ([]Address, error)

	// GetAddress retrieves a specific address
	//
	// Possible errors:
	//   - ErrNotFound: if the address does not exist
	GetAddress(context.Context, id.External) (Address, error)

	// SaveAddresses adds a new address or updates an existing one in the database
	//
	// Possible errors:
	//   - ErrAlreadyExists: if the address already exists for a given target
	SaveAddresses(context.Context, ...Address) ([]Address, error)

	// RemoveAddress deletes an address with the given id
	RemoveAddress(context.Context, id.External) error

	// ListPhones retrieves all the phones of a specific target id
	ListPhones(context.Context, id.External) ([]Phone, error)

	// GetPhone retrieves a specific phone
	//
	// Possible errors:
	//   - ErrNotFound: if the phone does not exist
	GetPhone(context.Context, id.External) (Phone, error)

	// SavePhones adds a new phone or updates an existing one in the database
	//
	// Possible errors:
	//   - ErrAlreadyExists: if the phone already exists for a given target
	SavePhones(context.Context, ...Address) ([]Phone, error)

	// RemovePhone deletes a phone with the given id
	RemovePhone(context.Context, id.External) error
}
