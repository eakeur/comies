package bill

import (
	"comies/core/billing/bill"
	"comies/io/data/postgres/conn"
	"context"
)

func (a actions) Create(ctx context.Context, b bill.Bill) error {
	const script = `
		insert into bills (
			id,
			reference_id,
			date,
			description,
		) values (
			$1, $2, $3, $4
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		b.ID,
		b.ReferenceID,
		b.Date,
		b.Name,
	)
	if err != nil {
		return err
	}

	return nil
}
