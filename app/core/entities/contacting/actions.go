package contacting

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {

	// ListAddresses retrieves all the addresses of a specific target id
	ListAddresses(ctx context.Context, target types.UID) ([]Address, error)

	// GetAddress retrieves a specific address
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the address does not exist
	GetAddress(ctx context.Context, target types.UID, addressID types.UID) (Address, error)

	// SaveAddresses adds a new address or updates an existing one
	SaveAddresses(ctx context.Context, target types.UID, addresses ...Address) ([]Address, error)

	// RemoveAddresses deletes an address with the given id or all addresses of a given target if no id is specified
	RemoveAddresses(ctx context.Context, target types.UID, ids ...types.UID) error

	// ListPhones retrieves all the phones of a specific target id
	ListPhones(ctx context.Context, target types.UID) ([]Phone, error)

	// GetPhone retrieves a specific phone
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the phone does not exist
	GetPhone(ctx context.Context, target types.UID, id types.UID) (Phone, error)

	// SavePhones adds a new phone or updates an existing one
	//
	// Possible errors:
	//   - fault.ErrAlreadyExists: if the phone already exists for a given target
	SavePhones(ctx context.Context, target types.UID, phones ...Phone) ([]Phone, error)

	// RemovePhones deletes a phone with the given id or all phones of a given target if no id is specified
	RemovePhones(ctx context.Context, target types.UID, ids ...types.UID) error
}
