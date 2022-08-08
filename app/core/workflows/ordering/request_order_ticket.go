package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) RequestOrderTicket(ctx context.Context) (types.ID, error) {
	o := order.Order{
		Status: order.InTheCartStatus,
	}

	w.id.Create(&o.ID)
	o, err := w.orders.Create(ctx, o)
	if err != nil {
		return 0, throw.Error(err)
	}

	return o.ID, nil
}
