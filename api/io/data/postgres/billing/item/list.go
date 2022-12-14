package item

import (
	"comies/core/billing/item"
	"comies/io/data/postgres/conn"
	"comies/io/data/postgres/query"
	"context"
)

func (a actions) List(ctx context.Context, filter item.Filter) ([]item.Item, error) {
	const script = `
		select
			i.id,
			i.bill_id,
			i.reference_id,
			i.credits,
			i.debts,
			i.description
		from
			bill_items i
		where
			%query%
	`

	q := query.NewQuery(script).
		Where(filter.BillID != 0, "i.bill_id= $%v", filter.BillID).
		Where(filter.ReferenceID != 0, "i.reference_id= $%v", filter.ReferenceID)

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, m item.Item) error {
			return scan(
				&m.ID,
				&m.BillID,
				&m.ReferenceID,
				&m.Credits,
				&m.Debts,
				&m.Description,
			)
		},
	)
}
