package order

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock
type Actions interface {
	CreateOrder(ctx context.Context, o Order) (Order, error)
	CreateItem(ctx context.Context, item Item) (Item, error)
	CreateContent(ctx context.Context, c ...Content) ([]Content, error)

	ListOrders(ctx context.Context, f Filter) ([]Order, int, error)
	ListItems(ctx context.Context, orderUID types.UID) ([]Item, error)
	ListContent(ctx context.Context, itemUID types.UID) ([]Content, error)

	GetOrder(ctx context.Context, uid types.UID) (Order, error)

	UpdateOrderStatus(ctx context.Context, uid types.UID, status Status) error
	UpdateOrderDeliveryMode(ctx context.Context, uid types.UID, deliverType DeliveryMode) error
	UpdateOrderAddressID(ctx context.Context, uid types.UID, addressID types.UID) error
	UpdateItemStatus(ctx context.Context, uid types.UID, status PreparationStatus) error
	UpdateContentStatus(ctx context.Context, uid types.UID, status PreparationStatus) error
	UpdateContentQuantity(ctx context.Context, uid types.UID, qt types.Quantity) error

	RemoveOrder(ctx context.Context, o Order) error
	RemoveItem(ctx context.Context, uid types.UID) error
	RemoveContent(ctx context.Context, uid types.UID) error
}
