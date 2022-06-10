package stock

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	CreateStock(ctx context.Context, st Stock) (Stock, error)
	UpdateStock(ctx context.Context, st Stock) error
	GetStockByID(ctx context.Context, resourceID types.ID) (Stock, error)
	RemoveStock(ctx context.Context, resourceID types.ID) error
	ListStocks(ctx context.Context) ([]Stock, int, error)
}
