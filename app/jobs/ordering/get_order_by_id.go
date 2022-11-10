package ordering

import (
	"comies/app/core/ordering/order"
	"comies/app/core/types"
	"context"
)

func (w jobs) GetOrderByID(ctx context.Context, id types.ID) (order.Order, error) {
	return w.orders.GetByID(ctx, id)
}
