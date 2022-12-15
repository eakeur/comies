package product

import (
	"comies/core/menu/product"
	"comies/io/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, p product.Product) error {
	const script = `
		insert into products (
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
		) values (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)
	`

	if _, err := conn.ExecFromContext(ctx, script,
		p.ID,
		p.Code,
		p.Name,
		p.Type,
		p.CostPrice,
		p.SaleUnit,
		p.MinimumSale,
		p.MinimumQuantity,
		p.MaximumQuantity,
		p.Location,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if conn.IsCode(pgErr, conn.DuplicateError) && conn.IsConstraint(pgErr, conn.ProductCodeUK) {
				return product.ErrCodeAlreadyExists
			}
		}

		return err
	}

	return nil
}
