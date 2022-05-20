package item

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"gomies/app/core/entities/item"
	"gomies/app/gateway/persistence/postgres"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
)

func (a actions) Create(ctx context.Context, i item.Item) (item.Item, error) {
	const script = `
		insert into items (
			id,
			order_id,
			status,
            price,
			product_id,
			quantity,
			observations,
			store_id
		) values (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`

	_, err := transaction.ExecFromContext(ctx, script,
		i.ID,
		i.OrderID,
		i.Status,
		i.Price,
		i.ProductID,
		i.Quantity,
		i.Observations,
		i.StoreID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.ItemOrderIDFK {
				return item.Item{}, fault.Wrap(fault.ErrNotFound).
					Describe("the order id provided seems to not exist").Params(map[string]interface{}{
					"order_id": i.OrderID.String(),
				})
			}
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.ItemIDPK {
				return item.Item{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the item id provided seems to already exist").Params(map[string]interface{}{
					"order_id": i.OrderID.String(), "item_id": i.ID.String(),
				})
			}
		}

		return item.Item{}, fault.Wrap(err)
	}

	return i, nil
}