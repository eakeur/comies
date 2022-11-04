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
	AddToOrder(ctx context.Context, i item.Item) (item.Item, error)
	PlaceOrder(ctx context.Context, o OrderConfirmation) (order.Order, error)
	CancelOrder(ctx context.Context, id types.ID) error

	ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error)
	ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error)
	CountUnfinishedOrders(ctx context.Context) (order.CountByStatus, error)

	SetOrderStatus(ctx context.Context, id types.ID, st status.Status) error
	SetItemStatus(ctx context.Context, itemID types.ID, st types.Status) error

	GetOrderByID(ctx context.Context, id types.ID) (order.Order, error)
	GetOrderByCustomerPhone(ctx context.Context, phone string) (order.Order, error)
}

type ProductPriceFetcher func(ctx context.Context, productID types.ID) (types.Currency, error)

type ProductDispatcher func(ctx context.Context, productID, agentID types.ID, quantity types.Quantity) error

type jobs struct {
	orders          order.Actions
	items           item.Actions
	statuses        status.Actions
	getPrice        ProductPriceFetcher
	createID        types.CreateID
	dispatchProduct ProductDispatcher
}

func NewJobs(
	orders order.Actions,
	items item.Actions,
	statuses status.Actions,
	createID types.CreateID,
	getPrice ProductPriceFetcher,
	dispatchProduct ProductDispatcher,
) Jobs {
	return jobs{
		statuses:        statuses,
		orders:          orders,
		items:           items,
		createID:        createID,
		getPrice:        getPrice,
		dispatchProduct: dispatchProduct,
	}
}
