package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/data/orders"
	"context"
)

func RequestOrderTicket(ctx context.Context) (id.ID, error) {
	o := order.Order{
		Status: order.InTheCartStatus,
	}

	id.Create(&o.ID)
	o, err := orders.Create(ctx, o)
	if err != nil {
		return 0, err
	}

	return o.ID, nil
}
