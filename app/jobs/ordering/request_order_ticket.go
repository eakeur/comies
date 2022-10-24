package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"context"
)

func (w jobs) RequestOrderTicket(ctx context.Context) (types.ID, error) {
	o := order.Order{
		Status: order.InTheCartStatus,
	}

	w.id.Create(&o.ID)
	o, err := w.orders.Create(ctx, o)
	if err != nil {
		return 0, err
	}

	return o.ID, nil
}