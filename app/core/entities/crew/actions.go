package crew

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// Get returns an operator with the parameterized id
	//
	// Possible errors:
	//   - ErrNotFound: if the operator does not exist
	Get(context.Context, id.External) (Operator, error)

	// GetWithOperatorAndStoreNick returns an operator with the parameterized
	// nickname and store nickname
	//
	// Possible errors:
	//   - ErrNotFound: if the operator does not exist
	GetWithOperatorAndStoreNick(context.Context, string, string) (Operator, error)

	// List retrieves categories respecting the parameterized filter
	List(context.Context, Filter) ([]Operator, error)

	// Create adds an operator to the database
	//
	// Possible errors:
	//   - ErrAlreadyExists: if the nick informed already belongs to another Operator
	Create(context.Context, Operator) (Operator, error)

	// Remove deletes the operator with the id informed
	//
	// Possible errors:
	//   - ErrHasDependants: if there are any dependency in this operator
	Remove(context.Context, id.External) error

	// Update updates the operator
	Update(context.Context, Operator) error
}
