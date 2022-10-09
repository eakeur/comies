package product

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, p product.Product) (product.Product, error) {
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

	if _, err := transaction.ExecFromContext(ctx, script,
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
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.ProductCodeUK {
				return product.Product{}, product.ErrCodeAlreadyExists
			}
		}

		return product.Product{}, err
	}

	return p, nil
}
