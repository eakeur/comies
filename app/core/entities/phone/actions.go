package phone

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, target types.ID) ([]Phone, error)
	GetByID(ctx context.Context, phoneID types.ID) (Phone, error)
	Save(ctx context.Context, phones Phone) (Phone, error)
	Remove(ctx context.Context, id types.ID) error
	RemoveAllByTarget(ctx context.Context, target types.ID) error
}
