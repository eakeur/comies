package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) RemoveReserved(ctx context.Context, agentID types.ID) error {
	const script = `
		delete from 
			movements m	
		where 
			m.agent_id = $1 and m.type = $2
	`

	cmd, err := transaction.ExecFromContext(ctx, script, agentID, movement.ReservedType)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return throw.ErrNotFound
	}

	return nil
}
