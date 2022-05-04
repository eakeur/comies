package product

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListProducts(ctx context.Context, productFilter Filter) ([]Product, int, error)
	GetProducts(ctx context.Context, key Key) (Product, error)
	GetProductSaleInfo(ctx context.Context, key Key) (Sale, error)
	CreateProduct(ctx context.Context, prd Product) (Product, error)
	UpdateProduct(ctx context.Context, prd Product) error
	RemoveProduct(ctx context.Context, key Key) error
	ListIngredients(ctx context.Context, productKey Key) ([]Ingredient, error)
	SaveIngredients(ctx context.Context, productKey Key, ingredients ...Ingredient) ([]Ingredient, error)
	RemoveIngredient(ctx context.Context, productKey Key, ingredientID types.ID) error
	RemoveAllIngredients(ctx context.Context, productKey Key) error
}
