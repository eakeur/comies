package order

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock
type Actions interface {
	CreateOrder(ctx context.Context, o Order) (Order, error)
	CreateItem(ctx context.Context, item Item) (Item, error)
	CreateContent(ctx context.Context, c ...Content) ([]Content, error)

	ListOrders(ctx context.Context, f Filter) ([]Order, int, error)
	ListItems(ctx context.Context, orderUID types.ID) ([]Item, error)
	ListContent(ctx context.Context, itemUID types.ID) ([]Content, error)

	GetOrder(ctx context.Context, id types.ID) (Order, error)

	UpdateOrderStatus(ctx context.Context, id types.ID, status Status) error
	UpdateOrderDeliveryMode(ctx context.Context, id types.ID, deliverType DeliveryMode) error
	UpdateOrderAddressID(ctx context.Context, id types.ID, addressID types.ID) error
	UpdateItemStatus(ctx context.Context, id types.ID, status PreparationStatus) error
	UpdateContentStatus(ctx context.Context, id types.ID, status PreparationStatus) error
	UpdateContentQuantity(ctx context.Context, id types.ID, qt types.Quantity) error

	RemoveOrder(ctx context.Context, o Order) error
	RemoveItem(ctx context.Context, id types.ID) error
	RemoveContent(ctx context.Context, id types.ID) error
}
