package order

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
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
			return "", throw.Error(throw.ErrNotFound).
				Describe("the order id provided seems to not exist").Params(map[string]interface{}{
				"order_id": orderID,
			})
		}
		return "", throw.Error(err)
	}

	return o, nil
}
