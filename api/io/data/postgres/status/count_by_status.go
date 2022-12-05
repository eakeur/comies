package status

import (
	"comies/core/ordering/status"
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"context"
)

func (a actions) CountByStatus(ctx context.Context) (status.CountByStatus, error) {
	const script = `
		select
			ls.value,
			count(ls.order_id)
		from latest_statuses ls
		group by ls.value
	`

	rows, err := conn.QueryFromContext(ctx, script)
	if err != nil {
		return nil, err
	}

	c := make(status.CountByStatus, 0)
	for rows.Next() {
		var status types.Status
		var count types.Quantity
		if err := rows.Scan(&status, &count); err != nil {
			return nil, err
		}

		c[status] = count
	}

	return c, nil
}
