package ordering

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/data/orders"
	"context"
)

func CancelOrder(ctx context.Context, id id.ID) error {

	o, err := orders.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := ordering.CheckIfOrderIsCancelable(o); err != nil {
		return nil
	}

	return orders.UpdateFlow(ctx, ordering.NewFlow(id, ordering.CanceledOrderStatus))

}
