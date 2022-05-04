package stock

import (
	"context"
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListMovements(ctx context.Context, filter Filter) ([]Movement, int, error)
	SaveMovements(ctx context.Context, movement Movement) (Movement, error)
	UpdateReserved(ctx context.Context, agentID types.ID) error
	RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error
	RemoveReserved(ctx context.Context, agentID types.ID) error
	ArchiveMovements(ctx context.Context, filter Filter) error
	ComputeStock(ctx context.Context, filter Filter) (types.Quantity, error)
	CreateStock(ctx context.Context, st Stock) (Stock, error)
	UpdateStock(ctx context.Context, st Stock) error
	GetStockByID(ctx context.Context, resourceID types.ID) (Stock, error)
	RemoveStock(ctx context.Context, resourceID types.ID) error
	ListStocks(ctx context.Context, filter listing.Filter) ([]Stock, int, error)
}
