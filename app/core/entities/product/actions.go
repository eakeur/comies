package product

import (
	"context"
	"gomies/app/core/types/id"
)

// AdditionalDataToConsider is a flag type to point out which product related data an operation must consider
type AdditionalDataToConsider int

const (
	// Stock is a flag that points out that a function must consider, alongside the product, its stock information properties
	Stock AdditionalDataToConsider = iota

	// Sale is a flag that points out that a function must consider, alongside the product, its sale information properties
	Sale AdditionalDataToConsider = iota

	// All is a flag that points out that a function must consider, alongside the product, all its information properties
	All AdditionalDataToConsider = iota
)

type Actions interface {

	// List searches for products that respect the passed filter. If no products are found, an empty array of Product is returned
	List(context.Context, Filter) ([]Product, error)

	// Get returns a Product identified with the id provided. The flags inform which additional data
	// has to be retrieved as well. It can return an error if no product or additional data is found
	Get(context.Context, id.External, ...AdditionalDataToConsider) (Product, error)

	// Create saves a Product in the repository. It can return an error if a product with the same code
	// already exists or if any required information is nil
	Create(context.Context, Product) (Product, error)

	// Update updates the name, code and category of a Product with a specific ID. It returns an error if the product do
	// not exist or if any required information is nil
	Update(context.Context, Product) error

	// Remove deactivates a Product in the repository. It won't be ever retrieved by List and Get functions as long as
	// a manager changes its state. After a certain amount of time, the product will be effectively removed from the repository.
	// It will return an error if there is no product with this ID.
	Remove(context.Context, id.External) error
}
