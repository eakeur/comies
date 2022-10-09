package order

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
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
		return err
	}

	if cmd.RowsAffected() != 1 {
		return throw.ErrNotFound
	}

	return nil
}
