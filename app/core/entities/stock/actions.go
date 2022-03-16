package stock

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// ComputeStock computes the quantity stored of a specific resource
	//
	// Possible errors:
	//   - ErrMustHaveTargetID: if the filter property TargetID was not set
	ComputeStock(context.Context, Filter) (Actual, error)

	// GetMovement retrieves a stock movement with their details
	//
	// Possible errors:
	//   - ErrNotFound: if the movement does not exist
	GetMovement(context.Context, id.External) (Movement, error)

	// ListMovements retrieves stock movements for a specific resource
	ListMovements(context.Context, Filter) ([]Movement, error)

	// AddToStock adds a stock movement for a specific resource
	AddToStock(context.Context, Movement) (Movement, error)

	// RemoveFromStock removes a stock movement from a resource stock
	RemoveFromStock(context.Context, id.External) error
}
