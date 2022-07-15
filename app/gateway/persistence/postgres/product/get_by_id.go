package product

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
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
			p.sale_price,
			p.sale_unit,
			p.minimum_sale
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
			return product.Product{}, throw.Error(throw.ErrNotFound).
				Describe("the product id provided seems to not exist").Params(map[string]interface{}{
				"id": id,
			})
		}
		return product.Product{}, throw.Error(err)
	}

	return p, nil
}
