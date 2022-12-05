package ordering

import (
	"comies/core/ordering/order"
	"comies/core/types"
	"context"
)

func (w jobs) GetOrderByID(ctx context.Context, id types.ID) (order.Order, error) {
	return w.orders.GetByID(ctx, id)
}
