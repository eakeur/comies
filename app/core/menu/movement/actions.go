package movement

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Remove(ctx context.Context, id types.ID) error
	List(ctx context.Context, filter Filter) ([]Movement, error)
	Create(ctx context.Context, m Movement) error
	Balance(ctx context.Context, filter Filter) (types.Quantity, error)
}
