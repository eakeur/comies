package products

import (
	"comies/app/core/id"
	"comies/app/core/product"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetSaleInfoByID(ctx context.Context, productID id.ID) (product.Sale, error) {
	const script = `
		select
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale
		from
			products p
		where
			p.id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, productID)
	if err != nil {
		return product.Sale{}, err
	}

	var p product.Sale
	if err := row.Scan(
		&p.CostPrice,
		&p.SalePrice,
		&p.SaleUnit,
		&p.MinimumSale,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Sale{}, product.ErrNotFound
		}
		return product.Sale{}, err
	}

	return p, nil
}
