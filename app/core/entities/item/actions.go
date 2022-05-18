package item

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock
type Actions interface {
	Create(ctx context.Context, item Item) (Item, error)
	List(ctx context.Context, orderID types.ID) ([]Item, error)
	SetStatus(ctx context.Context, itemID types.ID, status Status) error
	Remove(ctx context.Context, itemID types.ID) error
}
