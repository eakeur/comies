package item

import (
	"comies/app/core/types"
	"comies/app/io/data/postgres/conn"
	"context"
)

func (a actions) SetObservation(ctx context.Context, itemID types.ID, obs string) error {
	const script = `update items set observations = $1 where id = $2`

	cmd, err := conn.ExecFromContext(ctx, script, obs, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() <= 0 {
		return types.ErrNotFound
	}

	return nil
}
