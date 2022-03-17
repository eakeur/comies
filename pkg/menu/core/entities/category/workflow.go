package category

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	// Save creates a new category or updates an existing one
	//
	// Possible errors
	//   - session.ErrNoSession: if there is no session in this context
	//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrAlreadyExists: if the category already exists and the "overwrite" flag was not set
	//   - ErrInvalidCode: if the code is invalid
	//   - ErrInvalidName: if the name is invalid
	Save(ctx context.Context, cat Category, flag ...types.WritingFlag) (Category, error)

	// List searches all categories with the given filter
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
	List(ctx context.Context, categoryFilter Filter) ([]Category, error)

	// Get retrieves a category with this categoryKey
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrNotFound: if the category does not exist
	Get(ctx context.Context, categoryKey Key) (Category, error)

	// Remove deletes a category
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
	Remove(ctx context.Context, categoryID Key) error
}
