package stock

import (
	"comies/app/sdk/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Create(ctx context.Context, st Stock) (Stock, error)
	Update(ctx context.Context, st Stock) error
	GetByID(ctx context.Context, resourceID types.ID) (Stock, error)
	Remove(ctx context.Context, resourceID types.ID) error
}
