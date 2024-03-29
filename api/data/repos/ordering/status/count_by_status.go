package status

import (
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) CountByStatus(ctx context.Context, status types.Status) (types.Quantity, error) {
	const script = `
		select count(ls.order_id) from latest_statuses ls where ls.value = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, status)
	if err != nil {
		return 0, err
	}

	var count types.Quantity
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
