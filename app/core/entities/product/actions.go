package product

import (
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListProducts(ctx context.Context, productFilter Filter) ([]Product, error)
	GetProducts(ctx context.Context, key Key) (Product, error)
	GetProductSaleInfo(ctx context.Context, key Key) (Sale, error)
	CreateProduct(ctx context.Context, prd Product) (Product, error)
	UpdateProduct(ctx context.Context, prd Product) error
	RemoveProduct(ctx context.Context, key Key) error
}
