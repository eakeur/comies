package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RequestOrderTicket(ctx context.Context, customerID types.ID) (types.ID, error) {

	if customerID.Empty() {
		return 0, fault.Wrap(fault.ErrMissingID)
	}

	o, err := w.orders.Create(ctx, order.Order{
		CustomerID: customerID,
		Status:     order.InTheCartStatus,
	})
	if err != nil {
		return 0, fault.Wrap(err)
	}

	return o.ID, nil
}
