package stock

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	// Compute calculates the quantity stored of a specific resource
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	Compute(ctx context.Context, filter Filter) (types.Quantity, error)

	// ComputeSome calculates all quantities for the resources specified by the resourcesIDs array
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	ComputeSome(ctx context.Context, filter Filter, resourcesIDs ...types.UID) ([]types.Quantity, error)

	// ListMovements retrieves movements (archived and not) of the resource identified with the resourceID
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
	SaveMovements(ctx context.Context, config Config, resourceID types.UID, movements ...Movement) (AdditionResult, error)

	// RemoveMovement deletes a movement from the stock of a given resource identified
	// by the parameterized resourceID. If the movementID is empty, it removes all movements
	// from the resource's stock
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	RemoveMovement(ctx context.Context, resourceID types.UID, movementID types.UID) error

	// ClosePeriod archives all movements from a given resource and period, blocking them from being deleted
	// and being counted on compute functions
	//
	// Possible errors:
	//   - ErrMissingResourceID: if the resourceID is invalid
	//   - ErrInvalidPeriod: if the dates are invalid or are not in a valid period
	ClosePeriod(ctx context.Context, filter Filter) error
}
