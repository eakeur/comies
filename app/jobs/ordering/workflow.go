package ordering

import (
	"comies/app/core/ordering/item"
	"comies/app/core/ordering/order"
	"comies/app/core/ordering/status"
	"comies/app/core/types"
	"context"
)

var _ Jobs = jobs{}

type Jobs interface {
	InitializeOrder(ctx context.Context) (types.ID, error)
	AddToOrder(ctx context.Context, i item.Item) (item.Item, error)
	PlaceOrder(ctx context.Context, o order.Order) (order.Order, error)

	ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error)
	ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error)
	CountUnfinishedOrders(ctx context.Context) (order.CountByStatus, error)

	SetOrderDeliveryType(ctx context.Context, id types.ID, deliveryType types.Type) error
	SetOrderStatus(ctx context.Context, id types.ID, st status.Status) error
	SetItemStatus(ctx context.Context, itemID types.ID, st types.Status) error
	SetItemObservation(ctx context.Context, itemID types.ID, obs string) error

	GetOrderByID(ctx context.Context, id types.ID) (order.Order, error)
	GetOrderByCustomerPhone(ctx context.Context, phone string) (order.Order, error)

	CancelOrder(ctx context.Context, id types.ID) error
}

type jobs struct {
	orders   order.Actions
	items    item.Actions
	statuses status.Actions
	createID types.CreateID
}

func NewJobs(
	orders order.Actions,
	items item.Actions,
	statuses status.Actions,
	createID types.CreateID,
) Jobs {
	return jobs{
		statuses: statuses,
		orders:   orders,
		items:    items,
		createID: createID,
	}
}
