package movement

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, filter Filter) ([]Movement, int, error)
	Save(ctx context.Context, movement Movement) (Movement, error)
	Remove(ctx context.Context, resourceID types.ID, movementID types.ID) error
	GetBalance(ctx context.Context, filter Filter) (types.Quantity, error)
	SetOutputStatus(ctx context.Context, agentID types.ID) error
	RemoveReserved(ctx context.Context, agentID types.ID) error
}
