package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
)

func (w workflow) ListOrders(ctx context.Context, f order.Filter) ([]order.Order, int, error) {
	list, count, err := w.orders.List(ctx, f)
	if err != nil {
		return nil, 0, fault.Wrap(err)
	}

	return list, count, nil
}
