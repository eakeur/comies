package movements

import (
	"comies/app/core/id"
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func SetOutputType(ctx context.Context, agentID id.ID) error {
	const script = `
		update movements set type = $1 where id = $2
	`

	cmd, err := conn.ExecFromContext(ctx, script, menu.OutputMovementType, agentID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
