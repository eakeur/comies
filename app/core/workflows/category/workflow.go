package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	CreateCategory(ctx context.Context, cat category.Category) (category.Category, error)
	UpdateCategory(ctx context.Context, cat category.Category) error
	ListCategories(ctx context.Context, categoryFilter category.Filter) ([]category.Category, error)
	GetCategory(ctx context.Context, categoryKey category.Key) (category.Category, error)
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
