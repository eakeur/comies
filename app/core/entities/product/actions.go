package product

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

// AdditionalDataToConsider is a flag type to point out which product related data an operation must consider
type AdditionalDataToConsider int

const (
	// Stock points out that stock information properties should be retrieved
	Stock AdditionalDataToConsider = iota

	// Sale points out that sale information properties should be retrieved
	Sale AdditionalDataToConsider = iota

	// All points out that all its information properties should be retrieved
	All AdditionalDataToConsider = iota
)

type Actions interface {

	// List searches for products that respect the passed filter.
	List(context.Context, Filter) ([]Product, error)

	// Get returns a Product identified with the id provided. The flags inform which additional data
	// has to be retrieved as well.
	//
	// Possible errors:
	//   - ErrNotFound: if there is no product found with such id
	Get(context.Context, id.External, ...AdditionalDataToConsider) (Product, error)

	// Create saves a Product in the repository.
	//
	// Possible errors:
	//   - ErrAlreadyExists: if there already is a product with the same code
	Create(context.Context, Product) (Product, error)

	// Update sets new properties for a Product with a specific ID.
	//
	// Possible errors:
	//   - ErrAlreadyExists: if there already is a product with the same code
	Update(context.Context, Product) error

	// Remove deletes a product from the database or deactivates it
	// if it has orders or stock movements
	Remove(context.Context, id.External) error
}
