package orders

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func SetDeliveryMode(ctx context.Context, id id.ID, deliverType order.DeliveryMode) error {
	const script = `
		update
			orders
		set
			delivery_mode = $1
		where
			id = $2
	`

	cmd, err := conn.ExecFromContext(ctx, script, deliverType, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
