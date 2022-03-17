package contacting

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {

	// ListAddresses retrieves all the addresses of a specific target id
	ListAddresses(ctx context.Context, target types.External) ([]Address, error)

	// GetAddress retrieves a specific address
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the address does not exist
	GetAddress(ctx context.Context, target types.External, addressID types.External) (Address, error)

	// SaveAddresses adds a new address or updates an existing one
	SaveAddresses(ctx context.Context, target types.External, addresses ...Address) ([]Address, error)

	// RemoveAddress deletes an address with the given id
	RemoveAddress(ctx context.Context, target types.External, id types.External) error

	// RemoveAllAddresses deletes all addresses for the target id
	RemoveAllAddresses(ctx context.Context, target types.External) error

	// ListPhones retrieves all the phones of a specific target id
	ListPhones(ctx context.Context, target types.External) ([]Phone, error)

	// GetPhone retrieves a specific phone
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the phone does not exist
	GetPhone(ctx context.Context, target types.External, id types.External) (Phone, error)

	// SavePhones adds a new phone or updates an existing one
	//
	// Possible errors:
	//   - fault.ErrAlreadyExists: if the phone already exists for a given target
	SavePhones(ctx context.Context, target types.External, phones ...Address) ([]Phone, error)

	// RemovePhone deletes a phone with the given id
	RemovePhone(ctx context.Context, target types.External, id types.External) error

	// RemoveAllPhones deletes all phones for a specific target
	RemoveAllPhones(ctx context.Context, target types.External) error
}
