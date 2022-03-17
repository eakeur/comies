package product

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {

	// List searches for products that respect the passed filter.
	List(ctx context.Context, productFilter Filter) ([]Product, error)

	// Get returns a Product identified with the id provided. The flags inform which additional data
	// has to be retrieved as well.
	//
	// Possible errors:
	//   - fault.ErrNotFound: if there is no product found with such id
	Get(ctx context.Context, key Key) (Product, error)

	// Save creates a Product in the repository or updates it if it already exists
	//
	// Possible errors:
	//   - fault.ErrAlreadyExists: if there already is a product with the same code and the "overwrite" flag was not set
	Save(ctx context.Context, prd Product, flag ...types.WritingFlag) (Product, error)

	// Remove deletes a product from the database or deactivates it
	// if it has orders or stock movements
	Remove(ctx context.Context, key Key) error
}
