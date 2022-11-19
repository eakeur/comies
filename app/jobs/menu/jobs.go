package menu

import (
	"comies/app/core/menu/ingredient"
	"comies/app/core/menu/movement"
	"comies/app/core/menu/price"
	"comies/app/core/menu/product"
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Jobs:WorkflowMock
type Jobs interface {
	CreateProduct(ctx context.Context, p ProductCreation) (types.ID, error)
	CreateMovement(ctx context.Context, m movement.Movement) (types.ID, error)
	CreateIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error)

	ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, error)
	ListProductsRunningOut(ctx context.Context) ([]MissingProduct, error)
	ListIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error)
	ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error)
	ListPrices(ctx context.Context, productID types.ID) ([]price.Price, error)

	GetProductByID(ctx context.Context, id types.ID) (product.Product, error)
	GetProductNameByID(ctx context.Context, id types.ID) (string, error)
	GetProductLatestPriceByID(ctx context.Context, id types.ID) (types.Currency, error)
	GetProductStockBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error)

	RemoveIngredient(ctx context.Context, productID, ingredientID types.ID) error
	RemoveMovement(ctx context.Context, id types.ID) error

	UpdateProduct(ctx context.Context, prd product.Product) error
	UpdateProductPrice(ctx context.Context, productID types.ID, val types.Currency) error
}

type jobs struct {
	products    product.Actions
	ingredients ingredient.Actions
	movements   movement.Actions
	prices      price.Actions
	createID    types.CreateID
}

var _ Jobs = jobs{}

func NewJobs(
	products product.Actions,
	ingredients ingredient.Actions,
	movements movement.Actions,
	prices price.Actions,
	createID types.CreateID,
) Jobs {
	return jobs{
		ingredients: ingredients,
		products:    products,
		movements:   movements,
		prices:      prices,
		createID:    createID,
	}
}
