package product

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	ApproveSale(ctx context.Context, req product.ApproveSaleRequest) error
	CreateIngredient(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error)
	CreateMovement(ctx context.Context, productID types.UID, mov Movement) (types.Quantity, error)
	CreateProduct(ctx context.Context, prd product.Product) (product.Product, error)
	GetProduct(ctx context.Context, key product.Key) (product.Product, error)
	ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error)
	RemoveIngredient(ctx context.Context, productKey product.Key, id types.UID) error
	RemoveProduct(ctx context.Context, key product.Key) error
	UpdateProduct(ctx context.Context, prd product.Product) error
}

var _ Workflow = workflow{}

func NewWorkflow(
	products product.Actions,
	categories category.Actions,
	stocks StockService,
) Workflow {
	return workflow{
		products:   products,
		categories: categories,
		stocks:     stocks,
	}
}

type workflow struct {
	products   product.Actions
	categories category.Actions
	stocks     StockService
}
