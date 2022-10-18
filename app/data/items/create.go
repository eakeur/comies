package items

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, i ordering.Item) error {
	const script = `
		insert into items (
			id,
			order_id,
			status,
            price,
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
		i.ProductID,
		i.Quantity,
		i.Observations,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.ItemOrderIDFK {
				return types.ErrNotFound
			}
			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.ItemIDPK {
				return types.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}
