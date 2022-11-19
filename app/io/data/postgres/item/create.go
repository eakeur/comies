package item

import (
	"comies/app/core/ordering/item"
	"comies/app/core/types"
	"comies/app/io/data/postgres/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, i item.Item) error {
	const script = `
		insert into items (
			id,
			order_id,
			status,
            value,
			product_id,
			quantity,
			observations
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		i.ID,
		i.OrderID,
		i.Status,
		i.Value,
		i.ProductID,
		i.Quantity,
		i.Observations,
	)
	if err != nil {
		if pgErr := new(pgconn.PgError); errors.As(err, &pgErr) {
			if conn.IsCode(pgErr, conn.NonexistentFK) && conn.IsConstraint(pgErr, conn.ItemOrderIDFK) {
				return types.ErrNotFound
			}
			if conn.IsCode(pgErr, conn.DuplicateError) && conn.IsConstraint(pgErr, conn.ItemIDPK) {
				return types.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}
