package status

import (
	"comies/core/ordering/status"
	"comies/data/conn"
	"context"
)

func (a actions) Update(ctx context.Context, s status.Status) error {
	const script = `
		insert into statuses (
			order_id,
			occurred_at,
			value
		) values (
			$1, $2, $3
		)
	`

	if _, err := conn.ExecFromContext(ctx, script,
		s.OrderID,
		s.OccurredAt,
		s.Value,
	); err != nil {
		// TODO: implement specific errors validation
		return err
	}

	return nil
}
