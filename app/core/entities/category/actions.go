package category

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// Get returns an account with the parameterized id
	//
	// Possible errors:
	//   - ErrNotFound: if the category does not exist
	Get(context.Context, id.External) (Category, error)

	// List retrieves categories respecting the parameterized filter
	List(context.Context, Filter) ([]Category, error)

	// Create adds a category to the database
	//
	// Possible errors:
	//   - ErrAlreadyExists: if the code informed already belongs to another category
	Create(context.Context, Category) (Category, error)

	// Remove deletes the category with the id informed
	//
	// Possible errors:
	//   - ErrHasDependants: if there are any products in this category
	Remove(context.Context, id.External) error

	// Update updates the category
	Update(context.Context, Category) error
}
