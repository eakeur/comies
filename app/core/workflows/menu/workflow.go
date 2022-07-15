package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/product"
	"comies/app/sdk/types"
	"context"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	ReserveProduct(ctx context.Context, reservation Reservation) (Reservation, error)
	UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error

	CreateProduct(ctx context.Context, prd product.Product) (product.Product, error)
	GetProduct(ctx context.Context, key product.Key) (product.Product, error)
	ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, error)
	RemoveProduct(ctx context.Context, key product.Key) error
	UpdateProduct(ctx context.Context, prd product.Product) error

	AddProductIngredient(ctx context.Context, ingredient ingredient.Ingredient) (ingredient.Ingredient, error)
	RemoveProductIngredient(ctx context.Context, id types.ID) error
	ListProductIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error)
}

var _ Workflow = workflow{}

func NewWorkflow(
	products product.Actions,
	ingredients ingredient.Actions,
	stocks StockService,
) Workflow {
	return workflow{
		ingredients: ingredients,
		products:    products,
		stocks:      stocks,
	}
}

type workflow struct {
	products    product.Actions
	ingredients ingredient.Actions
	stocks      StockService
}
