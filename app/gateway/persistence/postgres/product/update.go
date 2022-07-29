package product

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Update(ctx context.Context, prd product.Product) error {
	const script = `
		update 
			products
		set
			code = $1,
			name = $2,
			type = $3,
			cost_price = $4,
			sale_price = $5,
			sale_unit = $6,
			minimum_sale = $7,
			minimum_quantity = $8,
			maximum_quantity = $9,
			location = $10
		where 
			id = $11
	`

	cmd, err := transaction.ExecFromContext(ctx, script,
		prd.Code,
		prd.Name,
		prd.Type,
		prd.CostPrice,
		prd.SalePrice,
		prd.SaleUnit,
		prd.MinimumSale,
		prd.MinimumQuantity,
		prd.MaximumQuantity,
		prd.Location,
		prd.ID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.ProductCodeUK {
				return throw.Error(product.ErrCodeAlreadyExists).
					Describe("the product code provided seems to already exist").Params(map[string]interface{}{
					"code": prd.Code,
				})
			}
		}
		return throw.Error(err)
	}

	if cmd.RowsAffected() != 1 {
		return throw.Error(product.ErrNotFound)
	}

	return nil
}
