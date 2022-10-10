package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/core/workflows/menu"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
)

func CancelOrder(ctx context.Context, id id.ID) error {

	o, err := orders.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if o.Status >= order.PreparingStatus {
		return order.ErrAlreadyPreparing
	}

	items, err := items.List(ctx, o.ID)
	if err != nil {
		return err
	}

	for _, item := range items {
		err := menu.UpdateReservation(ctx, item.ID, false)
		if err != nil {
			return err
		}
	}

	return nil
}
