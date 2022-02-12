package product

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"gomies/app/core/types/id"
	"gomies/app/core/workspaces/transaction"
	"stonehenge/app/core/types/logger"
)

var _ Workflow = workflow{}

func NewWorkflow(
	pr product.Actions,
	st stock.Actions,
	ct category.Actions,
	tx transaction.Manager,
	lg logger.Logger,
	) Workflow {
	return workflow{
		products:     pr,
		stocks:       st,
		categories:   ct,
		transactions: tx,
		logger:       lg,
	}
}

type Workflow interface {
	Create(ctx context.Context, input CreateInput) (CreateOutput, error)
	List(ctx context.Context, filter product.Filter) ([]product.Product, error)
	Get(ctx context.Context, ext id.External) (product.Product, error)
	Remove(ctx context.Context, ext id.External) error
	RemoveFromStock(ctx context.Context, ext id.External) error
	AddToStock(ctx context.Context, movement stock.Movement) (stock.Movement, error)
	Update(ctx context.Context, prd product.Product) error
}

type workflow struct {
	products     product.Actions
	stocks       stock.Actions
	categories   category.Actions
	transactions transaction.Manager
	logger       logger.Logger
}
