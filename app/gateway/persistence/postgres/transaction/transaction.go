package transaction

import (
	"comies/app/core/throw"
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func FromContext(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(key).(pgx.Tx)
	if !ok {
		return nil, ErrNoTransaction
	}

	return tx, nil
}

func ExecFromContext(ctx context.Context, script string, parameters ...interface{}) (pgconn.CommandTag, error) {
	tx, err := FromContext(ctx)
	if err != nil {
		return nil, throw.Error(err)
	}

	cmd, err := tx.Exec(ctx, script, parameters...)
	if err != nil {
		return nil, throw.Error(err)
	}

	return cmd, err
}

func QueryRowFromContext(ctx context.Context, script string, parameters ...interface{}) (pgx.Row, error) {
	tx, err := FromContext(ctx)
	if err != nil {
		return nil, throw.Error(err)
	}

	return tx.QueryRow(ctx, script, parameters...), nil
}
