package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) CancelOrder(ctx context.Context, id types.ID) error {

	o, err := w.orders.GetByID(ctx, id)
	if err != nil {
		return throw.Error(err).Params(map[string]interface{}{
			"order_id": o.ID,
		})
	}

	if o.Status >= order.PreparingStatus {
		return throw.Error(order.ErrAlreadyPreparing).Params(map[string]interface{}{
			"order_status": o.Status,
		})
	}

	items, err := w.items.List(ctx, o.ID)
	if err != nil {
		return throw.Error(err).Params(map[string]interface{}{
			"order_id": o.ID,
		})
	}

	for _, item := range items {
		err := w.products.UpdateResources(ctx, item.ID, false)
		if err != nil {
			return throw.Error(err).Params(map[string]interface{}{
				"item_id": item.ID,
				"consume": false,
			})
		}
	}

	return nil
}
