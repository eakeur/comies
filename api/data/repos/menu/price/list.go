package price

import (
	"comies/core/menu/price"
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) List(ctx context.Context, targetID types.ID) ([]price.Price, error) {
	const script = `
		select
			id,
			target_id,
			date,
			value
		from prices
		where target_id = $1
	`

	rows, err := conn.QueryFromContext(ctx, script, targetID)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, p *price.Price) error {
			return scan(
				&p.ID,
				&p.TargetID,
				&p.Date,
				&p.Value,
			)
		},
	)
}
