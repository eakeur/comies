package order

import (
	"context"
	"gomies/pkg/sdk/types"
)

type Actions interface {
	CreateOrder(ctx context.Context, o Order) (Order, error)
	CreateItems(ctx context.Context, items ...Item) (Item, error)
	ListOrders(ctx context.Context) ([]Order, int, error)
	ListItems(ctx context.Context, orderUID types.UID) ([]Item, error)
	GetOrder(ctx context.Context, uid types.UID) (Order, error)
	UpdateItemStatus(ctx context.Context, uid types.UID, status ItemStatus) error
	UpdateOrderStatus(ctx context.Context, uid types.UID, status Status) error
}
