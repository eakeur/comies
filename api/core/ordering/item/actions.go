package item

import (
	"comies/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context, orderID types.ID) ([]Item, error)
	SetStatus(ctx context.Context, itemID types.ID, status types.Status) error
	SetQuantity(ctx context.Context, itemID types.ID, qt types.Quantity) error
	SetObservation(ctx context.Context, itemID types.ID, obs string) error
	Remove(ctx context.Context, itemID types.ID) error
}
