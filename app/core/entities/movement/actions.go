package movement

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListByProductID(ctx context.Context, productID types.ID, filter Filter) ([]Movement, error)
	Create(ctx context.Context, movement Movement) (Movement, error)
	Remove(ctx context.Context, movementID types.ID) error
	GetBalanceByProductID(ctx context.Context, productID types.ID, filter Filter) (types.Quantity, error)
	SetOutputType(ctx context.Context, agentID types.ID) error
	RemoveReserved(ctx context.Context, agentID types.ID) error
}
