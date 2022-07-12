package member

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
)

func (a actions) Remove(ctx context.Context, key member.Key) error {
	const script = `delete from members where id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, key.ID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
