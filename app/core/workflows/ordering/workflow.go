package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/sdk/id"
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
		SetItemStatus(ctx context.Context, itemID types.ID, st item.Status) error

		ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error)
		ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error)
		GetOrderByID(ctx context.Context, id types.ID) (order.Order, error)
		CancelOrder(ctx context.Context, id types.ID) error

		Channel(ctx context.Context) (chan OrderNotification, error)
	}

	workflow struct {
		products MenuService
		orders   order.Actions
		items    item.Actions
		id       id.Manager
		channel  map[types.ID]chan OrderNotification
	}
)

func NewWorkflow(orders order.Actions, items item.Actions, products MenuService, id id.Manager) Workflow {
	return workflow{
		channel:  make(map[types.ID]chan OrderNotification),
		products: products,
		orders:   orders,
		items:    items,
		id:       id,
	}
}
