package product

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetByID(ctx context.Context, id types.ID) (product.Product, error) {
	const script = `
		select
			id,
			code,
			name,
			type,
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale
		from
			products p
		where
			p.id = $1
	`

	row := a.db.QueryRow(ctx, script, id)

	var p product.Product
	if err := row.Scan(
		&p.ID,
		&p.Code,
		&p.Name,
		&p.Type,
		&p.CostPrice,
		&p.SalePrice,
		&p.SaleUnit,
		&p.MinimumSale,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Product{}, fault.Wrap(fault.ErrNotFound).
				Describe("the product id provided seems to not exist").Params(map[string]interface{}{
				"id": id,
			})
		}
		return product.Product{}, fault.Wrap(err)
	}

	return p, nil
}
