package orders

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func SetObservation(ctx context.Context, id id.ID, observation string) error {
	const script = `
		update
			orders
		set
			observations = $1
		where
			id = $2
	`

	cmd, err := conn.ExecFromContext(ctx, script, observation, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
