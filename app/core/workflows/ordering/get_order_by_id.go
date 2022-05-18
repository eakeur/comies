package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetOrderByID(ctx context.Context, id types.ID) (order.Order, error) {
	o, err := w.orders.GetByID(ctx, id)
	if err != nil {
		return order.Order{}, fault.Wrap(err)
	}

	return o, nil
}
