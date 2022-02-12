package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/fault"
)

func (w workflow) Update(ctx context.Context, prd product.Product) error {
	const operation = "Workflows.Product.Update"
	w.transactions.Begin(ctx)
	defer w.transactions.Rollback(ctx)
	if err := prd.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}
	w.transactions.Commit(ctx)
	return w.products.Update(ctx, prd)
}
