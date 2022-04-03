package stock

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// Compute calculates the quantity stored of a specific resource
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	Compute(ctx context.Context, filter Filter) (types.Quantity, error)

	// ComputeSome calculates the quantity stored of all resources parameterized
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	ComputeSome(ctx context.Context, filter Filter, resourceID ...types.UID) ([]types.Quantity, error)

	// ListMovements retrieves movements of the resource identified with the resourceID
	// filtering also by the properties set in the filter property
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	ListMovements(ctx context.Context, filter Filter) ([]Movement, error)

	// SaveMovements creates movements related to the stock of a given resource, which
	// is identified by the parametrized resourceID.
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	SaveMovements(ctx context.Context, movement ...Movement) ([]Movement, error)

	// RemoveMovement deletes a movement from the stock of a given resource identified
	// by the parameterized resourceID.
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	RemoveMovement(ctx context.Context, resourceID types.UID, movementID types.UID) error

	// RemoveAllMovements deletes all movements from the stock of a given resource identified
	// by the parameterized resourceID.
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	RemoveAllMovements(ctx context.Context, resourceID types.UID) error

	// ArchiveMovements moves all movements in a given period into a specific storage for consulting purposes only
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	ArchiveMovements(ctx context.Context, filter Filter) error
}
