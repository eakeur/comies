package store

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// Get returns a store with the parameterized id
	//
	// Possible errors:
	//   - ErrNotFound: if the store does not exist
	Get(context.Context, id.External) (Store, error)

	// GetWithNick returns a store with the parameterized nickname
	//
	// Possible errors:
	//   - ErrNotFound: if the store does not exist
	GetWithNick(context.Context, string) (Store, error)

	// List retrieves categories respecting the parameterized filter
	List(context.Context, Filter) ([]Store, error)

	// Create adds a store to the database
	//
	// Possible errors:
	//   - ErrAlreadyExists: if the nick informed already belongs to another Store
	Create(context.Context, Store) (Store, error)

	// Remove deletes the store with the id informed
	//
	// Possible errors:
	//   - ErrHasDependants: if there are any dependency in this store
	Remove(context.Context, id.External) error

	// Update updates the store
	Update(context.Context, Store) error
}
