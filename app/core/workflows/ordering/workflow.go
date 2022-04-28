package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/pkg/sdk/types"
)

type (
	Workflow interface {
		CreateOrder(ctx context.Context, o order.Order) (order.Order, error)
		CreateItem(ctx context.Context, i order.Item) (order.Item, error)

		UpdateOrderDeliveryMode(ctx context.Context, uid types.UID, deliveryMode order.DeliveryMode) error
		UpdateOrderStatus(ctx context.Context, uid types.UID, st order.Status) error
		UpdateOrderAddressID(ctx context.Context, uid types.UID, addressID types.UID) error

		UpdateItemStatus(ctx context.Context, uid types.UID, status order.PreparationStatus) error

		UpdateContentStatus(ctx context.Context, uid types.UID, status order.PreparationStatus) error
		UpdateContentQuantity(ctx context.Context, uid types.UID, qt types.Quantity) error

		ListOrders(ctx context.Context, f order.Filter) ([]order.Order, int, error)
		GetOrder(ctx context.Context, uid types.UID) (order.Order, error)
		CancelOrder(ctx context.Context, uid types.UID) error
	}
)
