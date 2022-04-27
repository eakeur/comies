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

	ListOrders(ctx context.Context) ([]Order, int, error)
	ListItems(ctx context.Context, orderUID types.UID) ([]Item, error)
	ListContent(ctx context.Context, itemUID types.UID) ([]Content, error)

	GetOrder(ctx context.Context, uid types.UID) (Order, error)

	UpdateItemStatus(ctx context.Context, uid types.UID, status ItemStatus) error
	UpdateOrderStatus(ctx context.Context, uid types.UID, status Status) error

	UpdateOrder(ctx context.Context, o Order) error
	UpdateItem(ctx context.Context, i Item) error
	UpdateContent(ctx context.Context, c Content) error

	RemoveOrder(ctx context.Context, o Order) error
	RemoveItem(ctx context.Context, uid types.UID) error
	RemoveContent(ctx context.Context, uid types.UID) error
}
