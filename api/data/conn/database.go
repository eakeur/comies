package conn

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Scan = func(...interface{}) error

type Executer interface {
	Exec(ctx context.Context, script string, parameters ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func WithContext(ctx context.Context, e Executer) context.Context {
	return context.WithValue(ctx, key, e)
}

func FromContext(ctx context.Context) (Executer, error) {
	conn, ok := ctx.Value(key).(Executer)
	if !ok {
		return nil, ErrNoConnection
	}

	return conn, nil
}

func TXFromContext(ctx context.Context) (pgx.Tx, error) {
	conn, ok := ctx.Value(key).(pgx.Tx)
	if !ok {
		return nil, ErrNoConnection
	}

	return conn, nil
}

func ExecFromContext(ctx context.Context, script string, parameters ...interface{}) (pgconn.CommandTag, error) {
	tx, err := FromContext(ctx)
	if err != nil {
		return nil, err
	}

	cmd, err := tx.Exec(ctx, script, parameters...)
	if err != nil {
		return nil, err
	}

	return cmd, err
}

func QueryRowFromContext(ctx context.Context, script string, parameters ...interface{}) (pgx.Row, error) {
	tx, err := FromContext(ctx)
	if err != nil {
		return nil, err
	}

	return tx.QueryRow(ctx, script, parameters...), nil
}

func QueryFromContext(ctx context.Context, script string, parameters ...interface{}) (pgx.Rows, error) {
	tx, err := FromContext(ctx)
	if err != nil {
		return nil, err
	}

	return tx.Query(ctx, script, parameters...)
}

func ScanRows[T any](r pgx.Rows, fn func(scan Scan, v *T) error) ([]T, error) {
	arr := make([]T, 0)

	for r.Next() {
		var v T

		if err := fn(r.Scan, &v); err != nil {
			return nil, err
		}

		arr = append(arr, v)
	}

	return arr, nil
}

func IsCode(err *pgconn.PgError, code string) bool {
	return err.Code == code
}

func IsConstraint(err *pgconn.PgError, constraint string) bool {
	return err.ConstraintName == constraint
}
