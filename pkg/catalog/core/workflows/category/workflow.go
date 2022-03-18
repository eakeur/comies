package category

import (
	"context"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {

	// SaveCategory creates a new category or updates an existing one
	//
	// Possible errors
	//   - session.ErrNoSession: if there is no session in this context
	//   - session.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrAlreadyExists: if the category already exists and the "overwrite" flag was not set
	//   - category.ErrInvalidCode: if the code is invalid
	//   - category.ErrInvalidName: if the name is invalid
	SaveCategory(ctx context.Context, cat category.Category, flag ...types.WritingFlag) (category.Category, error)

	// ListCategories searches all categories with the given filter
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - session.ErrNotAllowed: if the session owner is not allowed to perform this operation
	ListCategories(ctx context.Context, categoryFilter category.Filter) ([]category.Category, error)

	// GetCategory retrieves a category with this categoryKey
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - session.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrNotFound: if the category does not exist
	GetCategory(ctx context.Context, categoryKey category.Key) (category.Category, error)

	// RemoveCategory deletes a category
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - session.ErrNotAllowed: if the session owner is not allowed to perform this operation
	RemoveCategory(ctx context.Context, categoryID category.Key) error
}

var _ Workflow = workflow{}

func NewWorkflow(
	categories category.Actions,
) Workflow {
	return workflow{
		categories: categories,
	}
}

type workflow struct {
	categories category.Actions
}
