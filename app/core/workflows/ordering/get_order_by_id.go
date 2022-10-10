package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/data/orders"
	"context"
)

func GetOrderByID(ctx context.Context, id id.ID) (order.Order, error) {
	o, err := orders.GetByID(ctx, id)
	if err != nil {
		return order.Order{}, err
	}

	return o, nil
}
