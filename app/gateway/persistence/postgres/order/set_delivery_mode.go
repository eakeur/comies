package order

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) SetDeliveryMode(ctx context.Context, id types.ID, deliverType order.DeliveryMode) error {
	const script = `
		update
			orders
		set
			delivery_mode = $1
		where
			id = $2
	`

	cmd, err := transaction.ExecFromContext(ctx, script, deliverType, id)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
