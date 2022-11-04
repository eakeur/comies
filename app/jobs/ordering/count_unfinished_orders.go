package ordering

import (
	"comies/app/core/ordering/order"
	"context"
)

func (w jobs) CountUnfinishedOrders(ctx context.Context) (order.CountByStatus, error) {
	return w.orders.CountUnfinished(ctx)
}
