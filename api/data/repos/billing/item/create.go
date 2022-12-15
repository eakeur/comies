package item

import (
	"comies/core/billing/item"
	"comies/data/conn"
	"context"
)

func (a actions) Create(ctx context.Context, i item.Item) error {
	const script = `
		insert into bill_items (
			id,
			bill_id,
			reference_id,
			credits,
			debts,
			description
		) values (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		i.ID,
		i.BillID,
		i.ReferenceID,
		i.Credits,
		i.Debts,
		i.Description,
	)
	if err != nil {
		return err
	}

	return nil
}
