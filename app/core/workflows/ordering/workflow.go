package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/types"
)

type (
	Workflow interface {
		CreateOrder(ctx context.Context, o order.Order) (order.Order, error)
		CreateItem(ctx context.Context, i order.Item) (order.Item, error)

		UpdateOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error
		UpdateOrderStatus(ctx context.Context, id types.ID, st order.Status) error
		UpdateOrderAddressID(ctx context.Context, id types.ID, addressID types.ID) error

		UpdateItemStatus(ctx context.Context, id types.ID, status order.PreparationStatus) error

		UpdateContentStatus(ctx context.Context, id types.ID, status order.PreparationStatus) error
		UpdateContentQuantity(ctx context.Context, id types.ID, qt types.Quantity) error

		ListOrders(ctx context.Context, f order.Filter) ([]order.Order, int, error)
		GetOrder(ctx context.Context, id types.ID) (order.Order, error)
		CancelOrder(ctx context.Context, id types.ID) error
	}
)
