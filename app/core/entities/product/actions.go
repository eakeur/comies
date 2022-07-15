package product

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, productFilter Filter) ([]Product, error)
	GetByID(ctx context.Context, id types.ID) (Product, error)
	GetByCode(ctx context.Context, code string) (Product, error)
	GetSaleInfoByID(ctx context.Context, productID types.ID) (Sale, error)
	Create(ctx context.Context, prd Product) (Product, error)
	Update(ctx context.Context, prd Product) error
	Remove(ctx context.Context, id types.ID) error
}
