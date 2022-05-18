package order

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock
type Actions interface {
	Create(ctx context.Context, o Order) (Order, error)
	List(ctx context.Context, f Filter) ([]Order, int, error)
	GetByID(ctx context.Context, id types.ID) (Order, error)
	UpdateStatus(ctx context.Context, id types.ID, status Status) error
	UpdateDeliveryMode(ctx context.Context, id types.ID, deliverType DeliveryMode) error
	UpdateAddressID(ctx context.Context, id types.ID, addressID types.ID) error
	Remove(ctx context.Context, o Order) error
}
