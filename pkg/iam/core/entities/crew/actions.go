package crew

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// Get returns an operator with the parameterized key
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the operator does not exist
	Get(ctx context.Context, key Key) (Member, error)

	// GetWithNicknames returns an operator with the parameterized
	// nickname and store nickname
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the operator does not exist
	GetWithNicknames(ctx context.Context, operatorNickname string, storeNickname string) (Member, error)

	// List retrieves categories respecting the parameterized filter
	List(ctx context.Context, operatorFilter Filter) ([]Member, error)

	// Save adds an operator or updates it
	//
	// Possible errors:
	//   - fault.ErrAlreadyExists: if the nick informed already belongs to another Member
	//  and the "overwrite" flag is not set
	Save(ctx context.Context, op Member, flag ...types.WritingFlag) (Member, error)

	// Remove deletes the operator with the id informed
	//
	// Possible errors:
	//   - fault.ErrHasDependants: if there are any dependency in this operator
	Remove(ctx context.Context, key Key) error
}
