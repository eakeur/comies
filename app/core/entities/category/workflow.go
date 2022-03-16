package category

import (
	"context"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	Create(context.Context, Category) (Category, error)
	List(context.Context, Filter) ([]Category, error)
	Get(context.Context, id.External) (Category, error)
	Remove(context.Context, id.External) error
	Update(context.Context, Category) error
}
