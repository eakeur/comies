package product

import (
	"comies/core/menu/product"
	"comies/io/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetByCode(ctx context.Context, code string) (product.Product, error) {
	const script = `
		select
			id,
			code,
			name,
			type,
			cost_price,
			sale_unit,
			minimum_sale,
			minimum_quantity,
			maximum_quantity,
			location
		from
			products p
		where
			p.code = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, code)
	if err != nil {
		return product.Product{}, err
	}

	var p product.Product
	if err := row.Scan(
		&p.ID,
		&p.Code,
		&p.Name,
		&p.Type,
		&p.CostPrice,
		&p.SaleUnit,
		&p.MinimumSale,
		&p.MinimumQuantity,
		&p.MaximumQuantity,
		&p.Location,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Product{}, product.ErrNotFound
		}
		return product.Product{}, err
	}

	return p, nil
}
