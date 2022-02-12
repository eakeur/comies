package product

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"gomies/app/core/managers/transaction"
	"gomies/app/core/types/id"
)

var _ Workflow = workflow{}

func NewWorkflow(
	pr product.Actions,
	st stock.Actions,
	ct category.Actions,
	tx transaction.Manager,
) Workflow {
	return workflow{
		products:     pr,
		stocks:       st,
		categories:   ct,
		transactions: tx,
	}
}

type Workflow interface {
	Create(context.Context, CreateInput) (product.Product, error)
	List(context.Context, product.Filter) ([]product.Product, error)
	Get(context.Context, id.External) (product.Product, error)
	Remove(context.Context, id.External) error
	RemoveFromStock(context.Context, id.External) error
	AddToStock(context.Context, stock.Movement) (stock.Movement, error)
	ListStock(context.Context, stock.Filter) ([]stock.Movement, error)
	Update(context.Context, product.Product) error
}

type workflow struct {
	products     product.Actions
	stocks       stock.Actions
	categories   category.Actions
	transactions transaction.Manager
}
