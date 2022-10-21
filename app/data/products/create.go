package products

import (
	"comies/app/core/product"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context) func(p product.Product) error {
	return func(p product.Product) error {
		const script = `
		insert into products (
			id,
			code,
			name,
			type,
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale,
			minimum_quantity,
			maximum_quantity,
			location
		) values (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
	`

		if _, err := conn.ExecFromContext(ctx, script,
			p.ID,
			p.Code,
			p.Name,
			p.Type,
			p.CostPrice,
			p.SalePrice,
			p.SaleUnit,
			p.MinimumSale,
			p.MinimumQuantity,
			p.MaximumQuantity,
			p.Location,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.ProductCodeUK {
					return product.ErrCodeAlreadyExists
				}
			}

			return err
		}

		return nil
	}
}
