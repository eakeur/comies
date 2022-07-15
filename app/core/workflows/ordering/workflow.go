package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/sdk/types"
	"context"
)

var _ Workflow = workflow{}

type (
	Workflow interface {
		RequestOrderTicket(ctx context.Context) (types.ID, error)
		Order(ctx context.Context, o OrderConfirmation) (order.Order, error)
		AddToOrder(ctx context.Context, i item.Item) (ItemAdditionResult, error)

		SetOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error
		SetOrderStatus(ctx context.Context, id types.ID, st order.Status) error

		ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error)
		GetOrderByID(ctx context.Context, id types.ID) (order.Order, error)
		CancelOrder(ctx context.Context, id types.ID) error
	}

	workflow struct {
		products MenuService
		orders   order.Actions
		items    item.Actions
	}
)

func NewWorkflow(orders order.Actions, items item.Actions, products MenuService) Workflow {
	return workflow{
		products: products,
		orders:   orders,
		items:    items,
	}
}
