package category

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type (
	Actions interface {
		// Get returns an account with the parameterized key
		//
		// Possible errors:
		//   - fault.ErrNotFound: if the category does not exist
		Get(ctx context.Context, categoryKey Key) (Category, error)

		// List retrieves categories respecting the given filter
		List(ctx context.Context, categoryFilter Filter) ([]Category, error)

		// Save creates a new category or updates it if it already exists
		//
		// Possible errors:
		//   - fault.ErrAlreadyExists: if the category already exists and the
		//  "overwrite" flag was not set
		Save(ctx context.Context, cat Category, flag ...types.WritingFlag) (Category, error)

		// Remove deletes the category with the id informed
		//
		// Possible errors:
		//   - fault.ErrResourceHasChildren: if there are any products in this category
		Remove(ctx context.Context, categoryID Key) error
	}
)
