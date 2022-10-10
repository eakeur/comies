package products

import (
	"comies/app/core/id"
	"comies/app/core/product"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetStockInfoByID(ctx context.Context, productID id.ID) (product.Stock, error) {
	const script = `
		select
			minimum_quantity,
			maximum_quantity,
			location
		from
			products p
		where
			p.id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, productID)
	if err != nil {
		return product.Stock{}, err
	}

	var p product.Stock
	if err := row.Scan(
		&p.MinimumQuantity,
		&p.MaximumQuantity,
		&p.Location,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Stock{}, product.ErrNotFound
		}
		return product.Stock{}, err
	}

	return p, nil
}
