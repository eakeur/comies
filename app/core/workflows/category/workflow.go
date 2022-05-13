package category

import (
	"context"
	category2 "gomies/app/core/entities/category"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	CreateCategory(ctx context.Context, cat category2.Category) (category2.Category, error)
	UpdateCategory(ctx context.Context, cat category2.Category) error
	ListCategories(ctx context.Context, categoryFilter category2.Filter) ([]category2.Category, int, error)
	GetCategory(ctx context.Context, categoryKey category2.Key) (category2.Category, error)
	RemoveCategory(ctx context.Context, categoryID category2.Key) error
}

var _ Workflow = workflow{}

func NewWorkflow(
	categories category2.Actions,
) Workflow {
	return workflow{
		categories: categories,
	}
}

type workflow struct {
	categories category2.Actions
}
