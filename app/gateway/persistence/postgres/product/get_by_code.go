package product

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
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
			sale_price,
			sale_unit,
			minimum_sale
		from
			products p
		where
			p.code = $1
	`

	row := a.db.QueryRow(ctx, script, code)

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
				Describe("the product code provided seems to not exist").Params(map[string]interface{}{
				"code": code,
			})
		}
		return product.Product{}, throw.Error(err)
	}

	return p, nil
}
