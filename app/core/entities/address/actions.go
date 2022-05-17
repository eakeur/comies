package address

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, targetID types.ID) ([]Address, error)
	GetByID(ctx context.Context, addressID types.ID) (Address, error)
	Save(ctx context.Context, address Address) (Address, error)
	Remove(ctx context.Context, id types.ID) error
	RemoveAllByTarget(ctx context.Context, target types.ID) error
}
