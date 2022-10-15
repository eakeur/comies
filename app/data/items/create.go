package items

import (
	"comies/app/core/item"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, i item.Item) error {
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
		i.Price,
		i.ProductID,
		i.Quantity,
		i.Observations,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.ItemOrderIDFK {
				return item.Item{}, types.ErrNotFound
			}
			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.ItemIDPK {
				return item.Item{}, types.ErrAlreadyExists
			}
		}

		return item.Item{}, err
	}

	return i, nil
}
