package ordering

import (
	"comies/core/ordering/item"
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/core/types"
	"comies/jobs/billing"
	"comies/jobs/menu"
	"context"
)

var _ Jobs = jobs{}

type Jobs interface {
	PlaceOrder(ctx context.Context, o Order) (order.Order, error)
	CancelOrder(ctx context.Context, id types.ID) error

	ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error)
	ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error)
	CountOrdersByStatus(ctx context.Context, status types.Status) (types.Quantity, error)

	SetOrderStatus(ctx context.Context, id types.ID, st status.Status) error
	SetItemStatus(ctx context.Context, itemID types.ID, st types.Status) error

	GetOrderByID(ctx context.Context, id types.ID) (order.Order, error)
	GetStatusByCustomerPhone(ctx context.Context, phone string) (Status, error)
}

type jobs struct {
	orders   order.Actions
	items    item.Actions
	statuses status.Actions
	createID types.CreateID

	menu    menu.Jobs
	billing billing.Jobs
}

type Deps struct {
	Orders    order.Actions
	Items     item.Actions
	Statuses  status.Actions
	IDCreator types.CreateID

	Menu    menu.Jobs
	Billing billing.Jobs
}

func NewJobs(deps Deps) Jobs {
	return jobs{
		statuses: deps.Statuses,
		orders:   deps.Orders,
		items:    deps.Items,
		createID: deps.IDCreator,

		menu:    deps.Menu,
		billing: deps.Billing,
	}
}
