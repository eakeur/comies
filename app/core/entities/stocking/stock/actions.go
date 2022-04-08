package stock

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Compute(ctx context.Context, filter Filter) (types.Quantity, error)
	ComputeSome(ctx context.Context, filter Filter, resourceID ...types.UID) ([]types.Quantity, error)
	ListMovements(ctx context.Context, filter Filter) ([]Movement, int, error)
	SaveMovements(ctx context.Context, movement ...Movement) ([]Movement, error)
	RemoveMovement(ctx context.Context, resourceID types.UID, movementID types.UID) error
	RemoveAllMovements(ctx context.Context, resourceID types.UID) error
	ArchiveMovements(ctx context.Context, filter Filter) error
}
