package ordering

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/types"
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
		products ProductService
		orders   order.Actions
		items    item.Actions
	}
)

func NewWorkflow(orders order.Actions, products ProductService) Workflow {
	return workflow{
		products: products,
		orders:   orders,
	}
}
