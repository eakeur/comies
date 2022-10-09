package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"context"
)

func (w workflow) GetOrderByID(ctx context.Context, id types.ID) (order.Order, error) {
	o, err := w.orders.GetByID(ctx, id)
	if err != nil {
		return order.Order{}, err
	}

	return o, nil
}
