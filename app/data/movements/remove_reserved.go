package movements

import (
	"comies/app/core/id"
	"comies/app/core/movement"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func RemoveReserved(ctx context.Context, agentID id.ID) error {
	const script = `
		delete from 
			movements m	
		where 
			m.agent_id = $1 and m.type = $2
	`

	cmd, err := conn.ExecFromContext(ctx, script, agentID, movement.ReservedType)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return types.ErrNotFound
	}

	return nil
}
