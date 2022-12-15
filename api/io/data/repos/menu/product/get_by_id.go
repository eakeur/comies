package product

import (
	"comies/core/menu/product"
	"comies/core/types"
	"comies/io/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetByID(ctx context.Context, id types.ID) (product.Product, error) {
	const script = `
		select
			p.id,
			p.code,
			p.name,
			p.type,
			p.cost_price,
			p.sale_unit,
			p.minimum_sale,
			p.minimum_quantity,
			p.maximum_quantity,
			p.location
		from
			products p
		where
			p.id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, id)
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
