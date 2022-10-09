package order

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"comies/app/core/types"
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
			return 0, throw.ErrNotFound
		}
		return 0, err
	}

	return o, nil
}
