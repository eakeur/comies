package product

import (
	"context"
	"gomies/app/core/entities/product"
)

func (w workflow) Update(ctx context.Context, prd product.Product) error {
	w.transactions.Begin(ctx)
	defer w.transactions.Rollback(ctx)
	if err := prd.Validate(); err != nil {
		return err
	}
	w.transactions.Commit(ctx)
	return w.products.Update(ctx, prd)
}
