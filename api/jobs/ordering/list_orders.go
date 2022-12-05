package ordering

import (
	"comies/core/ordering/order"
	"context"
)

func (w jobs) ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error) {
	return w.orders.List(ctx, f)
}
