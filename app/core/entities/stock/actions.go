package stock

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Compute(ctx context.Context, filter Filter) (types.Quantity, error)
	ComputeSome(ctx context.Context, filter Filter, resourceID ...types.ID) ([]types.Quantity, error)
	ListMovements(ctx context.Context, filter Filter) ([]Movement, int, error)
	SaveMovements(ctx context.Context, movement ...Movement) ([]Movement, error)
	RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error
	RemoveAllMovements(ctx context.Context, resourceID types.ID) error
	ArchiveMovements(ctx context.Context, filter Filter) error
}
