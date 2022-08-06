package product

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, productFilter Filter) ([]Product, error)
	ListRunningOut(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id types.ID) (Product, error)
	GetByCode(ctx context.Context, code string) (Product, error)
	GetNameByID(ctx context.Context, id types.ID) (string, error)
	GetSaleInfoByID(ctx context.Context, productID types.ID) (Sale, error)
	GetStockInfoByID(ctx context.Context, productID types.ID) (Stock, error)
	Create(ctx context.Context, prd Product) (Product, error)
	Update(ctx context.Context, prd Product) error
	Remove(ctx context.Context, id types.ID) error
}
