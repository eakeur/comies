package product

import (
	"context"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) Remove(ctx context.Context, ext product.Key) error {
	const operation = "Workflows.Product.Remove"

	_, err := session.DelegateSessionProps(ctx, operation, &ext.Store, nil)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	err = w.products.Remove(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return err

}
