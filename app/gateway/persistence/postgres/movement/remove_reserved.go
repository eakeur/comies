package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/fault"
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
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() < 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
