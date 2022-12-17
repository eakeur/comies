package order

import (
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) Remove(ctx context.Context, id types.ID) error {
	const script = `delete from orders where id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
