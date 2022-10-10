package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/data/orders"
	"context"
)

func SetOrderDeliveryMode(ctx context.Context, id id.ID, deliveryMode order.DeliveryMode) error {
	err := orders.SetDeliveryMode(ctx, id, deliveryMode)
	if err != nil {
		return err
	}

	return nil
}
