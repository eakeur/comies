package ordering

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/data/orders"
	"context"
)

func SetOrderDeliveryMode(ctx context.Context, id id.ID, deliveryType ordering.Type) error {
	if err := ordering.ValidateDeliveryType(deliveryType); err != nil {
		return err
	}

	return orders.SetDeliveryMode(ctx, id, deliveryType)
}
