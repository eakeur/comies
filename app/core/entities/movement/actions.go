package movement

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListByResourceID(ctx context.Context, resourceID types.ID, filter Filter) ([]Movement, error)
	Create(ctx context.Context, movement Movement) (Movement, error)
	Remove(ctx context.Context, movementID types.ID) error
	GetBalanceByResourceID(ctx context.Context, resourceID types.ID, filter Filter) (types.Quantity, error)
	SetOutputType(ctx context.Context, agentID types.ID) error
	RemoveReserved(ctx context.Context, agentID types.ID) error
}
