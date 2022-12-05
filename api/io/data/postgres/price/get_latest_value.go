package price

import (
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"context"
)

func (a actions) GetLatestValue(ctx context.Context, targetID types.ID) (types.Currency, error) {
	const script = `
		select
			p.value
		from
			prices p
		where
			p.target_id = $1
		order by
			p.date desc
		limit 1
	`

	row, err := conn.QueryRowFromContext(ctx, script, targetID)
	if err != nil {
		return 0, err
	}

	var value types.Currency
	if err := row.Scan(&value); err != nil {
		// TODO: implement specific error validations
		return 0, err
	}

	return value, nil
}
