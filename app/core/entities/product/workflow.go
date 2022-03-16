package product

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	Create(context.Context, Product) (Product, error)
	List(context.Context, Filter) ([]Product, error)
	Get(context.Context, id.External) (Product, error)
	Remove(context.Context, id.External) error
	RemoveFromStock(context.Context, id.External) error
	AddToStock(context.Context, stock.Movement) (stock.Movement, error)
	ListStock(context.Context, stock.Filter) ([]stock.Movement, error)
	Update(context.Context, Product) error
}
