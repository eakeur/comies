package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (a actions) RemoveReserved(ctx context.Context, agentID types.ID) error {
	const script = `
		delete from 
			movements m	
		where 
			m.agent_id = $1 and m.type = $2
	`

	cmd, err := transaction.ExecFromContext(ctx, script, agentID, movement.ReservedMovement)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() < 1 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
