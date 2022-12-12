package item

import (
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"context"
)

func (a actions) SetStatus(ctx context.Context, itemID types.ID, status types.Status) error {
	const script = `update items set status = $1 where id = $2`

	cmd, err := conn.ExecFromContext(ctx, script, status, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() <= 0 {
		return types.ErrNotFound
	}

	return nil
}
