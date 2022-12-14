package menu

import (
	"comies/core/menu/ingredient"
	"comies/core/menu/movement"
	"comies/core/menu/price"
	"comies/core/menu/product"
	"comies/core/types"
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

	DispatchProduct(ctx context.Context, d Dispatcher) error
}

type jobs struct {
	products    product.Actions
	ingredients ingredient.Actions
	movements   movement.Actions
	prices      price.Actions
	createID    types.CreateID
}

type Deps struct {
	Products    product.Actions
	Ingredients ingredient.Actions
	Movements   movement.Actions
	Prices      price.Actions
	IDCreator   types.CreateID
}

var _ Jobs = jobs{}

func NewJobs(deps Deps) Jobs {
	return jobs{
		ingredients: deps.Ingredients,
		products:    deps.Products,
		movements:   deps.Movements,
		prices:      deps.Prices,
		createID:    deps.IDCreator,
	}
}
