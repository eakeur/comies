package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RequestOrderTicket(ctx context.Context) (types.ID, error) {
	o, err := w.orders.Create(ctx, order.Order{
		Status: order.InTheCartStatus,
	})
	if err != nil {
		return 0, fault.Wrap(err)
	}

	return o.ID, nil
}
