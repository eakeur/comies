package product

import (
	"comies/app/core/entities/product"
	"comies/app/core/types"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetStockInfoByID(ctx context.Context, productID types.ID) (product.Stock, error) {
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

	row := a.db.QueryRow(ctx, script, productID)

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
