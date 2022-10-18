package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/orders"
	"context"
)

func SetOrderStatus(ctx context.Context, id types.ID, st ordering.Status) error {
	if err := ordering.ValidateOrderStatus(st); err != nil {
		return err
	}

	err := orders.UpdateFlow(ctx, ordering.NewFlow(id, st))
	if err != nil {
		return err
	}

	defer sendch(id, changeStatusPath, st)

	return nil
}
