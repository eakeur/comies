package product

import (
	"comies/app/core/entities/product"
	"comies/app/core/throw"
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
			return product.Stock{}, throw.Error(product.ErrNotFound).
				Describe("the product id provided seems to not exist").Params(map[string]interface{}{
				"productID": productID,
			})
		}
		return product.Stock{}, throw.Error(err)
	}

	return p, nil
}
