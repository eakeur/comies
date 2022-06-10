package order

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) GetStatus(ctx context.Context, orderID types.ID) (order.Status, error) {
	const script = `
		select
			s.status
		from
			orders_statuses s
		where
			s.order_id = $1
	`

	row := a.db.QueryRow(ctx, script, orderID)

	var o order.Status
	if err := row.Scan(
		&o,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fault.Wrap(fault.ErrNotFound).
				Describe("the order id provided seems to not exist").Params(map[string]interface{}{
				"order_id": orderID,
			})
		}
		return "", fault.Wrap(err)
	}

	return o, nil
}
