package bill

import (
	"comies/core/billing/bill"
	"comies/io/data/postgres/conn"
	"comies/io/data/postgres/query"
	"context"
)

func (a actions) List(ctx context.Context, filter bill.Filter) ([]bill.Bill, error) {
	const script = `
		select
			b.id,
			b.reference_id,
			b.date,
			b.name
		from
			bills b
		where
			%query%
	`

	q := query.NewQuery(script).
		Where(!filter.Period.Start.IsZero(), "b.date >= $%v", filter.Period.Start).And().
		Where(!filter.Period.End.IsZero(), "b.date <= $%v", filter.Period.End).And().
		Where(filter.ReferenceID != 0, "b.reference_id= $%v", filter.ReferenceID)

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, m bill.Bill) error {
			return scan(
				&m.ID,
				&m.ReferenceID,
				&m.Date,
				&m.Name,
			)
		},
	)
}
