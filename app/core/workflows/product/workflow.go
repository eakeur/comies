package product

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/ingredient"
	product2 "gomies/app/core/entities/product"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	ReserveProduct(ctx context.Context, reservation Reservation) (Reservation, error)
	UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error

	CreateProduct(ctx context.Context, prd product2.Product) (product2.Product, error)
	GetProduct(ctx context.Context, key product2.Key) (product2.Product, error)
	ListProducts(ctx context.Context, productFilter product2.Filter) ([]product2.Product, int, error)
	RemoveProduct(ctx context.Context, key product2.Key) error
	UpdateProduct(ctx context.Context, prd product2.Product) error

	AddProductIngredient(ctx context.Context, ingredient ingredient.Ingredient) (ingredient.Ingredient, error)
	RemoveProductIngredient(ctx context.Context, id types.ID) error
	ListProductIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error)
}

var _ Workflow = workflow{}

func NewWorkflow(
	products product2.Actions,
	categories category.Actions,
	ingredients ingredient.Actions,
	stocks StockService,
) Workflow {
	return workflow{
		ingredients: ingredients,
		products:    products,
		categories:  categories,
		stocks:      stocks,
	}
}

type workflow struct {
	products    product2.Actions
	categories  category.Actions
	ingredients ingredient.Actions
	stocks      StockService
}
