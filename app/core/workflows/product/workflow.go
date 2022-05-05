package product

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	ReserveProduct(ctx context.Context, reservation Reservation) (Reservation, error)
	UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error
	CreateIngredient(ctx context.Context, productKey product.Key, input IngredientInput) (product.Ingredient, error)
	CreateProduct(ctx context.Context, prd product.Product) (product.Product, error)
	GetProduct(ctx context.Context, key product.Key) (product.Product, error)
	ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error)
	RemoveIngredient(ctx context.Context, productKey product.Key, id types.ID) error
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
