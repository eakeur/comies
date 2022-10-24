package ordering

import (
	"comies/app/core/ordering/order"
	"comies/app/core/types"
	"context"
)

func (w jobs) GetOrderByID(ctx context.Context, id types.ID) (order.Order, error) {
	o, err := w.orders.GetByID(ctx, id)
	if err != nil {
		return order.Order{}, err
	}

	return o, nil
}
