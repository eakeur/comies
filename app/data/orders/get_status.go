package orders

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetStatus(ctx context.Context, orderID types.ID) (ordering.Status, error) {
	const script = `select s.status from orders_statuses s where s.order_id = $1`

	row, err := conn.QueryRowFromContext(ctx, script, orderID)
	if err != nil {
		return 0, err
	}

	var o ordering.Status
	if err := row.Scan(&o); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, types.ErrNotFound
		}

		return 0, err
	}

	return o, nil
}
