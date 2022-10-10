package movements

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func Remove(ctx context.Context, movementID id.ID) error {
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
