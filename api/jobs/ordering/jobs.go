package ordering

import (
	"comies/core/ordering/item"
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/core/types"
	"context"
)

var _ Jobs = jobs{}

type Jobs interface {
	PlaceOrder(ctx context.Context, o Order) (order.Order, error)
	CancelOrder(ctx context.Context, id types.ID) error

	ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error)
	ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error)
	CountUnfinishedOrders(ctx context.Context) (status.CountByStatus, error)

	SetOrderStatus(ctx context.Context, id types.ID, st status.Status) error
	SetItemStatus(ctx context.Context, itemID types.ID, st types.Status) error

	GetOrderByID(ctx context.Context, id types.ID) (order.Order, error)
	GetStatusByCustomerPhone(ctx context.Context, phone string) (Status, error)
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
