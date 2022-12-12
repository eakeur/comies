package movement

import (
	"comies/core/finance/movement"
	"comies/io/data/postgres/conn"
	"context"
)

func (a actions) Create(ctx context.Context, mov movement.Movement) error {
	const script = `
		insert into financial_movements (
			id,
			item_id,
			agent_id,
			type,
			date,
			description,
			price
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		mov.ID,
		mov.ItemID,
		mov.AgentID,
		mov.Type,
		mov.Date,
		mov.Description,
		mov.Price,
	)
	if err != nil {
		return err
	}

	return nil
}
