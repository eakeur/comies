package product

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"gomies/app/core/entities/movement"
	"gomies/app/core/entities/product"
	"gomies/app/gateway/persistence/postgres"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
)

func (a actions) Update(ctx context.Context, prd product.Product) error {
	const script = `
		update 
			products
		set
			code = $1
			name = $2
			type = $3
			cost_price = $4
			sale_price = $5
			sale_unit = $6
			minimum_sale = $7
		where 
			id = $9
	`

	cmd, err := transaction.ExecFromContext(ctx, script, movement.OutputMovement,
		prd.Code,
		prd.Name,
		prd.Type,
		prd.CostPrice,
		prd.SalePrice,
		prd.SaleUnit,
		prd.MinimumSale,
		prd.ID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.ProductCodeUK {
				return fault.Wrap(fault.ErrAlreadyExists).
					Describe("the product code provided seems to already exist").Params(map[string]interface{}{
					"code": prd.Code,
				})
			}
		}
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
