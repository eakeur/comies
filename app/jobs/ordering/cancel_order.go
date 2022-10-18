package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/orders"
	"context"
)

func CancelOrder(ctx context.Context, id types.ID) error {

	o, err := orders.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := ordering.CheckIfOrderIsCancelable(o); err != nil {
		return nil
	}

	return orders.UpdateFlow(ctx, ordering.NewFlow(id, ordering.CanceledOrderStatus))

}
