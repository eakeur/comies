package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemoveIngredient(ctx context.Context, productKey product.Key, id types.External) error {
	const operation = "Workflows.Product.RemoveIngredient"

	_, err := session.DelegateSessionProps(ctx, operation, &productKey.Store, nil)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	err = w.products.RemoveIngredient(ctx, productKey, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}
	return nil
}
