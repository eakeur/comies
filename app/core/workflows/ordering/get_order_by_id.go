package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) GetOrderByID(ctx context.Context, id types.ID) (order.Order, error) {
	o, err := w.orders.GetByID(ctx, id)
	if err != nil {
		return order.Order{}, fault.Wrap(err)
	}

	return o, nil
}
