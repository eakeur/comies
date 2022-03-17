package product

import (
	"context"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {

	// Save saves a product or updates an existing one
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrAlreadyExists: if the product already exists and the "overwrite" flag was not set
	//   - ErrInvalidCode: if the code is invalid
	//   - ErrInvalidName: if the name is invalid
	//   - ErrInvalidSalePrice: if the sal price is 0
	Save(ctx context.Context, prd Product, flag ...types.WritingFlag) (Product, error)

	// List retrieves all products respecting a given filter
	List(ctx context.Context, productFilter Filter) ([]Product, error)

	// Get retrieves a product respecting its key
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the product is not found
	Get(ctx context.Context, key Key) (Product, error)

	// Remove deletes a product or deactivates it if there is some children
	Remove(ctx context.Context, key Key) error

	// RemoveFromStock deletes a movement
	RemoveFromStock(ctx context.Context, productID types.External, movementID types.External) error

	// AddToStock adds a stock movement to the stock of a specific product
	//
	// Possible errors
	//   - fault.ErrNotFound: if the product does not exist
	//   - stocking.ErrStockAlreadyFull: if the stock is already full
	//   - stocking.ErrMustHaveTargetID: if the targetID for the movement was not set
	AddToStock(ctx context.Context, productID types.External, mov stock.Movement) (stock.Movement, error)

	// ListStock retrieves all movements for a specific target and filter
	ListStock(ctx context.Context, productID types.External, stockFilter stock.Filter) ([]stock.Movement, error)
}
