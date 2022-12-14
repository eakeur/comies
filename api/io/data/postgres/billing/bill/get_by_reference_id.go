package bill

import (
	"comies/core/billing/bill"
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetByReferenceID(ctx context.Context, id types.ID) (bill.Bill, error) {
	const script = `
		select
			b.id,
			b.reference_id,
			b.date,
			b.name
		from
			bill s
		where
			b.reference_id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, id)
	if err != nil {
		return bill.Bill{}, err
	}

	var b bill.Bill
	if err := row.Scan(
		&b.ID,
		&b.ReferenceID,
		&b.Name,
		&b.ReferenceID,
		&b.Date,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return bill.Bill{}, types.ErrNotFound
		}
		return bill.Bill{}, err
	}

	return b, nil
}
