package price

import (
	"comies/core/menu/price"
	"comies/data/conn"
	"context"
)

func (a actions) Create(ctx context.Context, p price.Price) error {
	const script = `
		insert into prices (
			id,
			target_id,
			date,
			value
		) values (
			$1, $2, $3, $4
		)
	`

	if _, err := conn.ExecFromContext(ctx, script,
		p.ID,
		p.TargetID,
		p.Date,
		p.Value,
	); err != nil {
		// TODO: implement specific errors validation
		return err
	}

	return nil
}
