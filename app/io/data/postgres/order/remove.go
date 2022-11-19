package order

import (
	"comies/app/core/types"
	"comies/app/io/data/postgres/conn"
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
