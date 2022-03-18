package product

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {

	// List searches for products that respect the passed filter.
	List(ctx context.Context, productFilter Filter) ([]Product, error)

	// Get returns a Product identified with the id provided.
	//
	// Possible errors:
	//   - fault.ErrNotFound: if there is no product found with such id
	Get(ctx context.Context, key Key) (Product, error)

	// GetProductSaleInfo returns a product's sale info identified with the id provided.
	//
	// Possible errors:
	//   - fault.ErrNotFound: if there is no product found with such id
	GetProductSaleInfo(ctx context.Context, key Key) (Sale, error)

	// GetProductStockInfo returns a product's stock info identified with the id provided.
	//
	// Possible errors:
	//   - fault.ErrNotFound: if there is no product found with such id
	GetProductStockInfo(ctx context.Context, key Key) (Stock, error)

	// Save creates a Product in the repository or updates it if it already exists
	//
	// Possible errors:
	//   - fault.ErrAlreadyExists: if there already is a product with the same code and the "overwrite" flag was not set
	Save(ctx context.Context, prd Product, flag ...types.WritingFlag) (Product, error)

	// Remove deletes a product from the database or deactivates it
	// if it has orders or stock movements
	Remove(ctx context.Context, key Key) error

	// ListIngredients retrieves all ingredients of a given product
	//
	// Possible errors:
	//   - fault.ErrNotFound: if there is no product found with such id
	ListIngredients(ctx context.Context, productKey Key) ([]Ingredient, error)

	// SaveIngredients creates or updates all ingredients passed in
	SaveIngredients(ctx context.Context, productKey Key, ingredients ...Ingredient) ([]Ingredient, error)

	// RemoveIngredient removes all ingredients of a given product
	RemoveIngredient(ctx context.Context, productKey Key, ingredientID types.External) error

	// RemoveAllIngredients removes all ingredients of a given product
	RemoveAllIngredients(ctx context.Context, productKey Key) error
}
