package orders

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"time"
)

func SetPlacementDate(ctx context.Context, id id.ID, date time.Time) error {
	const script = `
		update
			orders
		set
			placed_at = $1
		where
			id = $2
	`

	cmd, err := conn.ExecFromContext(ctx, script, date, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
