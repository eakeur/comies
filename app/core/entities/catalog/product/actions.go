package product

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListProducts(ctx context.Context, productFilter Filter) ([]Product, int, error)
	GetProducts(ctx context.Context, key Key) (Product, error)
	GetProductSaleInfo(ctx context.Context, key Key) (Sale, error)
	GetProductStockInfo(ctx context.Context, key Key) (Stock, error)
	CreateProduct(ctx context.Context, prd Product) (Product, error)
	UpdateProduct(ctx context.Context, prd Product) error
	RemoveProduct(ctx context.Context, key Key) error
	ListIngredients(ctx context.Context, productKey Key) ([]Ingredient, error)
	SaveIngredients(ctx context.Context, productKey Key, ingredients ...Ingredient) ([]Ingredient, error)
	RemoveIngredient(ctx context.Context, productKey Key, ingredientID types.UID) error
	RemoveAllIngredients(ctx context.Context, productKey Key) error
}
