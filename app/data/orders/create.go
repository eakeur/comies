package orders

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, o ordering.Order) error {
	const script = `
		insert into orders (
			id,
			placed_at,
			delivery_mode,
			observations,
			customer_name,
			customer_address,
			customer_phone
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	if _, err := conn.ExecFromContext(ctx, script,
		o.ID,
		o.PlacedAt,
		o.DeliveryType,
		o.Observations,
		o.Customer.Name,
		o.Customer.Address,
		o.Customer.Phone,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.OrderIDPK {
				return types.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}
