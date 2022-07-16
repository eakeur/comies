package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/product"
	"comies/app/sdk/id"
	"comies/app/sdk/types"
	"context"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	ReserveProduct(ctx context.Context, reservation Reservation) (Reservation, error)
	UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error

	CreateProduct(ctx context.Context, prd product.Product) (product.Product, error)
	GetProductByID(ctx context.Context, id types.ID) (product.Product, error)
	GetProductByCode(ctx context.Context, code string) (product.Product, error)
	GetProductNameByID(ctx context.Context, id types.ID) (string, error)
	ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, error)
	RemoveProduct(ctx context.Context, id types.ID) error
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
	id id.Manager,
) Workflow {
	return workflow{
		ingredients: ingredients,
		products:    products,
		stocks:      stocks,
		id:          id,
	}
}

type workflow struct {
	products    product.Actions
	ingredients ingredient.Actions
	stocks      StockService
	id          id.Manager
}
