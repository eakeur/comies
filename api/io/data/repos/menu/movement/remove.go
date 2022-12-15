package movement

import (
	"comies/core/types"
	"comies/io/data/conn"
	"context"
)

func (a actions) Remove(ctx context.Context, movementID types.ID) error {
	const script = `delete from movements m where m.id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, movementID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
