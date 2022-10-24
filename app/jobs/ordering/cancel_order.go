package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"context"
)

func (w jobs) CancelOrder(ctx context.Context, id types.ID) error {

	o, err := w.orders.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if o.Status >= order.PreparingStatus {
		return order.ErrAlreadyPreparing
	}

	items, err := w.items.List(ctx, o.ID)
	if err != nil {
		return err
	}

	for _, item := range items {
		err := w.products.UpdateReservation(ctx, item.ID, false)
		if err != nil {
			return err
		}
	}

	return nil
}
