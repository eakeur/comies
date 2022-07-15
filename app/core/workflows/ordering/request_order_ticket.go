package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RequestOrderTicket(ctx context.Context) (types.ID, error) {
	o, err := w.orders.Create(ctx, order.Order{
		Status: order.InTheCartStatus,
	})
	if err != nil {
		return 0, throw.Error(err)
	}

	return o.ID, nil
}
