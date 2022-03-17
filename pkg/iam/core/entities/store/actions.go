package store

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// Get returns a store with the parameterized key
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the store does not exist
	Get(ctx context.Context, storeID Key) (Store, error)

	// List retrieves categories respecting the parameterized filter
	List(ctx context.Context, storeFilter Filter) ([]Store, error)

	// Save adds a store to the database or updates an exiting one
	//
	// Possible errors:
	//   - fault.ErrAlreadyExists: if the nick informed already belongs
	//  to another store and the "overwrite" flag is not set
	Save(ctx context.Context, st Store) (Store, error)

	// Remove deletes the store with the key informed
	//
	// Possible errors:
	//   - fault.ErrResourceHasChildren: if there are any dependency in this store
	Remove(ctx context.Context, key Key) error

	// ListPreferences returns all preference of a store
	// If one want a specific module or key, just pass the key on the key parameter
	// Possible errors:
	//   - fault.ErrNotFound: if the store does not exist
	ListPreferences(ctx context.Context, storeKey Key, modules ...string) (types.Preferences, error)

	// SavePreferences saves all preference of a store
	// Possible errors:
	//   - fault.ErrNotFound: if the store does not exist
	SavePreferences(ctx context.Context, storeKey Key, pref types.Preferences) (types.Preferences, error)
}
