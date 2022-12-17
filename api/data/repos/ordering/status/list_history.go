package status

import (
	"comies/core/ordering/status"
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) ListHistory(ctx context.Context, orderID types.ID) ([]status.Status, error) {
	const script = `
		select
			s.order_id,
			s.occurred_at,
			s.value
		from statuses
		where order_id = $1
	`

	rows, err := conn.QueryFromContext(ctx, script, orderID)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, s *status.Status) error {
			return scan(
				&s.OrderID,
				&s.OccurredAt,
				&s.Value,
			)
		},
	)
}
